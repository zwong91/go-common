package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "region_test",
    srcs = ["region_test.go"],
    embed = [":region"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-tag/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "region",
    srcs = ["region.go"],
    importpath = "go-common/app/interface/main/app-tag/dao/region",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-tag/conf",
        "//app/interface/main/app-tag/model/region",
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
