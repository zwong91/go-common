package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "plugin_test",
    srcs = ["plugin_test.go"],
    embed = [":plugin"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-resource/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "plugin",
    srcs = ["plugin.go"],
    importpath = "go-common/app/interface/main/app-resource/dao/plugin",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-resource/conf",
        "//app/interface/main/app-resource/model/plugin",
        "//library/database/sql",
        "//library/log",
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
