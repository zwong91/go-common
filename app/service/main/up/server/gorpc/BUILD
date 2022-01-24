package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "gorpc_test",
    srcs = ["rpcserver_test.go"],
    embed = [":gorpc"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/up/model",
        "//library/net/rpc",
        "//library/time",
        "//vendor/github.com/davecgh/go-spew/spew",
    ],
)

go_library(
    name = "gorpc",
    srcs = ["rpcserver.go"],
    importpath = "go-common/app/service/main/up/server/gorpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/up/conf",
        "//app/service/main/up/model",
        "//app/service/main/up/service",
        "//library/ecode",
        "//library/log",
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
