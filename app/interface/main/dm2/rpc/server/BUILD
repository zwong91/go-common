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
        "//app/interface/main/dm2/conf",
        "//app/interface/main/dm2/model",
        "//app/interface/main/dm2/service",
        "//library/net/rpc",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "server",
    srcs = ["rpc.go"],
    importpath = "go-common/app/interface/main/dm2/rpc/server",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/dm2/conf",
        "//app/interface/main/dm2/model",
        "//app/interface/main/dm2/service",
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
