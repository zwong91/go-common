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
        "//app/service/main/tag/model",
        "//app/service/main/tag/rpc/client",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "server",
    srcs = ["rpc.go"],
    importpath = "go-common/app/service/main/tag/rpc/server",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/tag/conf",
        "//app/service/main/tag/model",
        "//app/service/main/tag/service",
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
