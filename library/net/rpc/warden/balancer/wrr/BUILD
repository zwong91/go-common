package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "wrr_test",
    srcs = ["wrr_test.go"],
    embed = [":wrr"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/net/metadata",
        "//library/net/rpc/warden/metadata",
        "//vendor/google.golang.org/grpc/balancer",
        "//vendor/google.golang.org/grpc/resolver",
    ],
)

go_library(
    name = "wrr",
    srcs = ["wrr.go"],
    importpath = "go-common/library/net/rpc/warden/balancer/wrr",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//library/net/metadata",
        "//library/net/rpc/warden/metadata",
        "//library/stat/summary",
        "//vendor/google.golang.org/grpc",
        "//vendor/google.golang.org/grpc/balancer",
        "//vendor/google.golang.org/grpc/balancer/base",
        "//vendor/google.golang.org/grpc/codes",
        "//vendor/google.golang.org/grpc/metadata",
        "//vendor/google.golang.org/grpc/resolver",
        "//vendor/google.golang.org/grpc/status",
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
        "//library/net/rpc/warden/balancer/wrr/test:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
