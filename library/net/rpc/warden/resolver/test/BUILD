package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "test_test",
    srcs = ["resovler_test.go"],
    embed = [":test"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/conf/env",
        "//library/naming",
        "//library/net/netutil/breaker",
        "//library/net/rpc/warden",
        "//library/net/rpc/warden/balancer/wrr",
        "//library/net/rpc/warden/proto/testproto",
        "//library/net/rpc/warden/resolver",
        "//library/time",
        "//vendor/github.com/stretchr/testify/assert",
        "//vendor/google.golang.org/grpc",
    ],
)

go_library(
    name = "test",
    srcs = ["mockdiscovery.go"],
    importpath = "go-common/library/net/rpc/warden/resolver/test",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/conf/env",
        "//library/naming",
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
