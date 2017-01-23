# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# A Dockerfile for creating an Elasticsearch instance that is designed
# to work with Kubernetes logging. Inspired by the Dockerfile
# dockerfile/elasticsearch

FROM java:openjdk-7-jre
MAINTAINER Ajeesh "ajeesh@peek.com"
ENV DEBIAN_FRONTEND noninteractive
ENV NODE_MASTER true
ENV NODE_DATA true
ENV NODE_INGEST true
ENV HTTP_ENABLE true
ENV NETWORK_HOST _site_
ENV HTTP_CORS_ENABLE true
ENV HTTP_CORS_ALLOW_ORIGIN *
ENV NUMBER_OF_MASTERS 2
ENV ES_HEAP_SIZE 3g
ENV HTTP_PORT 9200
ENV TRANSPORT_PORT 9300

RUN apt-get update && \
    apt-get install -y curl && \
    apt-get clean

RUN cd / && \
    curl -O https://download.elastic.co/elasticsearch/elasticsearch/elasticsearch-1.5.2.tar.gz && \
    tar xf elasticsearch-1.5.2.tar.gz && \
    rm elasticsearch-1.5.2.tar.gz

RUN mkdir -p /elasticsearch-1.5.2/config/templates

COPY elasticsearch.yml /elasticsearch-1.5.2/config/elasticsearch.yml
COPY template-k8s-logstash.json /elasticsearch-1.5.2/config/templates/template-k8s-logstash.json
COPY run.sh /
COPY elasticsearch_logging_discovery /

COPY repository-s3-5.1.1.zip  /tmp/repository-s3-5.1.1.zip

VOLUME ["/data"]
EXPOSE 9200 9300


RUN /elasticsearch-1.5.2/bin/plugin   --install elastic/elasticsearch-cloud-aws/2.5.1  --url   file:///tmp/elasticsearch-cloud-aws-2.5.1.zip
CMD ["/run.sh"]
