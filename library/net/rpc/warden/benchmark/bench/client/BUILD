package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "client",
    embed = [":client_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "client_lib",
    srcs = ["client.go"],
    importpath = "go-common/library/net/rpc/warden/benchmark/bench/client",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/net/netutil/breaker",
        "//library/net/rpc/warden",
        "//library/net/rpc/warden/benchmark/bench/proto",
        "//library/time",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/montanaflynn/stats",
        "//vendor/golang.org/x/net/context",
        "//vendor/google.golang.org/grpc",
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
