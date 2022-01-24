package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "gorpc_test",
    srcs = ["rpc_test.go"],
    embed = [":gorpc"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//app/service/main/coin/model"],
)

go_library(
    name = "gorpc",
    srcs = ["rpc.go"],
    importpath = "go-common/app/service/main/coin/server/gorpc",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/coin/api",
        "//app/service/main/coin/conf",
        "//app/service/main/coin/model",
        "//app/service/main/coin/service",
        "//library/net/rpc",
        "//library/net/rpc/context",
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
