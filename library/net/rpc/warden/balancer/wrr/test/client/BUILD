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
    importpath = "go-common/library/net/rpc/warden/balancer/wrr/test/client",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/exp/feature",
        "//library/log",
        "//library/net/rpc/warden",
        "//library/net/rpc/warden/proto/testproto",
        "//library/net/rpc/warden/resolver",
        "//library/net/rpc/warden/resolver/direct",
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
