/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package webhook implements a generic HTTP webhook plugin.
package webhook

import (
	"fmt"
	"time"

	"k8s.io/client-go/pkg/api"
	apierrors "k8s.io/client-go/pkg/api/errors"
	"k8s.io/client-go/pkg/apimachinery/registered"
	"k8s.io/client-go/pkg/runtime"
	"k8s.io/client-go/pkg/runtime/schema"
	runtimeserializer "k8s.io/client-go/pkg/runtime/serializer"
	"k8s.io/client-go/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	_ "k8s.io/client-go/pkg/apis/authorization/install"
)

type GenericWebhook struct {
	RestClient     *rest.RESTClient
	initialBackoff time.Duration
}

// NewGenericWebhook creates a new GenericWebhook from the provided kubeconfig file.
func NewGenericWebhook(kubeConfigFile string, groupVersions []schema.GroupVersion, initialBackoff time.Duration) (*GenericWebhook, error) {
	for _, groupVersion := range groupVersions {
		if !registered.IsEnabledVersion(groupVersion) {
			return nil, fmt.Errorf("webhook plugin requires enabling extension resource: %s", groupVersion)
		}
	}

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.ExplicitPath = kubeConfigFile
	loader := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{})

	clientConfig, err := loader.ClientConfig()
	if err != nil {
		return nil, err
	}
	codec := api.Codecs.LegacyCodec(groupVersions...)
	clientConfig.ContentConfig.NegotiatedSerializer = runtimeserializer.NegotiatedSerializerWrapper(runtime.SerializerInfo{Serializer: codec})

	restClient, err := rest.UnversionedRESTClientFor(clientConfig)
	if err != nil {
		return nil, err
	}

	// TODO(ericchiang): Can we ensure remote service is reachable?

	return &GenericWebhook{restClient, initialBackoff}, nil
}

// WithExponentialBackoff will retry webhookFn() up to 5 times with exponentially increasing backoff when
// it returns an error for which apierrors.SuggestsClientDelay() or apierrors.IsInternalError() returns true.
func (g *GenericWebhook) WithExponentialBackoff(webhookFn func() rest.Result) rest.Result {
	var result rest.Result
	WithExponentialBackoff(g.initialBackoff, func() error {
		result = webhookFn()
		return result.Error()
	})
	return result
}

// WithExponentialBackoff will retry webhookFn() up to 5 times with exponentially increasing backoff when
// it returns an error for which apierrors.SuggestsClientDelay() or apierrors.IsInternalError() returns true.
func WithExponentialBackoff(initialBackoff time.Duration, webhookFn func() error) error {
	backoff := wait.Backoff{
		Duration: initialBackoff,
		Factor:   1.5,
		Jitter:   0.2,
		Steps:    5,
	}

	var err error
	wait.ExponentialBackoff(backoff, func() (bool, error) {
		err = webhookFn()
		if _, shouldRetry := apierrors.SuggestsClientDelay(err); shouldRetry {
			return false, nil
		}
		if apierrors.IsInternalError(err) {
			return false, nil
		}
		if err != nil {
			return false, err
		}
		return true, nil
	})
	return err
}
