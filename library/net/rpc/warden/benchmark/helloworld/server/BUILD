package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "server",
    embed = [":server_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "server_lib",
    srcs = ["greeter_server.go"],
    importpath = "go-common/library/net/rpc/warden/benchmark/helloworld/server",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/net/rpc/warden",
        "//library/net/rpc/warden/proto/testproto",
        "//library/time",
        "//vendor/github.com/prometheus/client_golang/prometheus/promhttp",
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
