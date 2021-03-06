package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = [
        "conversion.go",
        "defaults.go",
        "doc.go",
        "field_constants.go",
        "generate.go",
        "helpers.go",
        "mapper.go",
        "meta.go",
        "ref.go",
        "register.go",
        "resource_helpers.go",
        "types.go",
        "zz_generated.deepcopy.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//pkg/api/meta:go_default_library",
        "//pkg/api/resource:go_default_library",
        "//pkg/apis/meta/v1:go_default_library",
        "//pkg/conversion:go_default_library",
        "//pkg/fields:go_default_library",
        "//pkg/genericapiserver/api/request:go_default_library",
        "//pkg/labels:go_default_library",
        "//pkg/runtime:go_default_library",
        "//pkg/runtime/schema:go_default_library",
        "//pkg/runtime/serializer:go_default_library",
        "//pkg/selection:go_default_library",
        "//pkg/types:go_default_library",
        "//pkg/util/intstr:go_default_library",
        "//pkg/util/labels:go_default_library",
        "//pkg/util/rand:go_default_library",
        "//pkg/util/sets:go_default_library",
        "//pkg/util/uuid:go_default_library",
        "//pkg/util/validation/field:go_default_library",
        "//vendor:github.com/davecgh/go-spew/spew",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "generate_test.go",
        "helpers_test.go",
        "ref_test.go",
        "resource_helpers_test.go",
    ],
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//pkg/api/resource:go_default_library",
        "//pkg/apis/meta/v1:go_default_library",
        "//pkg/labels:go_default_library",
        "//pkg/runtime:go_default_library",
        "//pkg/runtime/schema:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "conversion_test.go",
        "copy_test.go",
        "deep_copy_test.go",
        "defaulting_test.go",
        "meta_test.go",
        "serialization_proto_test.go",
        "serialization_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//pkg/api:go_default_library",
        "//pkg/api/meta:go_default_library",
        "//pkg/api/testapi:go_default_library",
        "//pkg/api/testing:go_default_library",
        "//pkg/api/v1:go_default_library",
        "//pkg/apimachinery/registered:go_default_library",
        "//pkg/apis/batch/v2alpha1:go_default_library",
        "//pkg/apis/extensions:go_default_library",
        "//pkg/apis/extensions/v1beta1:go_default_library",
        "//pkg/apis/meta/v1:go_default_library",
        "//pkg/conversion:go_default_library",
        "//pkg/genericapiserver/api/request:go_default_library",
        "//pkg/runtime:go_default_library",
        "//pkg/runtime/schema:go_default_library",
        "//pkg/runtime/serializer/protobuf:go_default_library",
        "//pkg/runtime/serializer/streaming:go_default_library",
        "//pkg/types:go_default_library",
        "//pkg/util/diff:go_default_library",
        "//pkg/util/sets:go_default_library",
        "//pkg/util/uuid:go_default_library",
        "//pkg/watch:go_default_library",
        "//vendor:github.com/davecgh/go-spew/spew",
        "//vendor:github.com/gogo/protobuf/proto",
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:github.com/google/gofuzz",
        "//vendor:github.com/spf13/pflag",
        "//vendor:github.com/ugorji/go/codec",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//pkg/api/annotations:all-srcs",
        "//pkg/api/endpoints:all-srcs",
        "//pkg/api/errors:all-srcs",
        "//pkg/api/events:all-srcs",
        "//pkg/api/install:all-srcs",
        "//pkg/api/meta:all-srcs",
        "//pkg/api/pod:all-srcs",
        "//pkg/api/resource:all-srcs",
        "//pkg/api/rest:all-srcs",
        "//pkg/api/service:all-srcs",
        "//pkg/api/testapi:all-srcs",
        "//pkg/api/testing:all-srcs",
        "//pkg/api/unversioned:all-srcs",
        "//pkg/api/util:all-srcs",
        "//pkg/api/v1:all-srcs",
        "//pkg/api/validation:all-srcs",
    ],
    tags = ["automanaged"],
)
