package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "direct_test",
    srcs = ["direct_test.go"],
    embed = [":direct"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/net/netutil/breaker",
        "//library/net/rpc/warden",
        "//library/net/rpc/warden/proto/testproto",
        "//library/net/rpc/warden/resolver",
        "//library/time",
    ],
)

go_library(
    name = "direct",
    srcs = ["direct.go"],
    importpath = "go-common/library/net/rpc/warden/resolver/direct",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/conf/env",
        "//library/naming",
        "//library/net/rpc/warden/resolver",
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
