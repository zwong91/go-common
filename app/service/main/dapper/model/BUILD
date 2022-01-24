package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "model",
    srcs = [
        "response.go",
        "span.go",
        "spanpoint.go",
        "util.go",
    ],
    importpath = "go-common/app/service/main/dapper/model",
    tags = ["manual"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/net/trace/proto",
        "//vendor/github.com/golang/protobuf/proto",
        "//vendor/github.com/golang/protobuf/ptypes/duration",
        "//vendor/github.com/golang/protobuf/ptypes/timestamp",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
