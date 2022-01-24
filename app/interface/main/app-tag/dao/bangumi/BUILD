package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "bangumi_test",
    srcs = ["bangumi_test.go"],
    embed = [":bangumi"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-tag/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "bangumi",
    srcs = ["bangumi.go"],
    importpath = "go-common/app/interface/main/app-tag/dao/bangumi",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-tag/conf",
        "//app/interface/main/app-tag/model/bangumi",
        "//library/log",
        "//library/net/http/blademaster",
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
