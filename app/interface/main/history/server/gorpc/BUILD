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
    deps = [
        "//app/interface/main/history/model",
        "//app/interface/main/history/rpc/client",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "gorpc",
    srcs = ["rpc.go"],
    importpath = "go-common/app/interface/main/history/server/gorpc",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/history/conf",
        "//app/interface/main/history/model",
        "//app/interface/main/history/service",
        "//library/ecode",
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
