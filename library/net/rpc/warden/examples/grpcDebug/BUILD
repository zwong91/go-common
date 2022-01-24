package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_library(
    name = "grpcDebug_lib",
    srcs = ["client.go"],
    importpath = "go-common/library/net/rpc/warden/examples/grpcDebug",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/jsonpb",
        "//vendor/google.golang.org/grpc",
        "//vendor/google.golang.org/grpc/credentials",
        "//vendor/google.golang.org/grpc/encoding",
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

go_binary(
    name = "grpcDebug",
    embed = [":grpcDebug_lib"],
    tags = ["automanaged"],
)
