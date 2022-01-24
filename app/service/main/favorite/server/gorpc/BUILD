package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "gorpc_test",
    srcs = ["server_test.go"],
    embed = [":gorpc"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/favorite/api/gorpc",
        "//app/service/main/favorite/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "gorpc",
    srcs = ["server.go"],
    importpath = "go-common/app/service/main/favorite/server/gorpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/favorite/conf",
        "//app/service/main/favorite/model",
        "//app/service/main/favorite/service",
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
