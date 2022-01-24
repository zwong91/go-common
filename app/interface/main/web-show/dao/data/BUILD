package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "data_test",
    srcs = ["data_test.go"],
    embed = [":data"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/web-show/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "data",
    srcs = ["data.go"],
    importpath = "go-common/app/interface/main/web-show/dao/data",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/web-show/conf",
        "//library/log",
        "//library/net/http/blademaster",
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
