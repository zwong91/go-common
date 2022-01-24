package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "server_test",
    srcs = ["rpc_test.go"],
    embed = [":server"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tag/conf",
        "//app/interface/main/tag/model",
        "//app/interface/main/tag/service",
        "//library/log",
    ],
)

go_library(
    name = "server",
    srcs = [
        "channel.go",
        "rpc.go",
        "rpc2.go",
    ],
    importpath = "go-common/app/interface/main/tag/rpc/server",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tag/conf",
        "//app/interface/main/tag/model",
        "//app/interface/main/tag/service",
        "//library/ecode",
        "//library/log",
        "//library/net/rpc",
        "//library/net/rpc/context",
        "//library/xstr",
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
