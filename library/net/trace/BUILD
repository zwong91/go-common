load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

package(default_visibility = ["//visibility:public"])

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
        "//library/net/trace/proto:all-srcs",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "trace",
    srcs = glob(
        ["*.go"],
        exclude = ["*_test.go"],
    ),
    importpath = "go-common/library/net/trace",
    tags = ["manual"],
    deps = [
        "//library/conf/dsn",
        "//library/conf/env",
        "//library/net/metadata",
        "//library/net/trace/proto",
        "//library/time",
        "//vendor/github.com/golang/protobuf/proto",
        "//vendor/github.com/golang/protobuf/ptypes/duration",
        "//vendor/github.com/golang/protobuf/ptypes/timestamp",
        "//vendor/github.com/pkg/errors",
        "//vendor/google.golang.org/grpc/metadata",
    ],
)

go_test(
    name = "trace_test",
    srcs = [
        "config_test.go",
        "context_test.go",
        "dapper_test.go",
        "marshal_test.go",
        "report_test.go",
        "sample_test.go",
        "span_test.go",
        "tag_test.go",
        "util_test.go",
    ],
    embed = [":trace"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/stretchr/testify/assert",
        "//vendor/google.golang.org/grpc/metadata",
    ],
)
