package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "recommend_test",
    srcs = ["recommend_test.go"],
    embed = [":recommend"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-tag/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "recommend",
    srcs = ["recommend.go"],
    importpath = "go-common/app/interface/main/app-tag/dao/recommend",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-tag/conf",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/time",
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
