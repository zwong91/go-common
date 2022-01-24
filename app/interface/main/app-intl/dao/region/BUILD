package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "region_test",
    srcs = ["dao_test.go"],
    embed = [":region"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-intl/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "region",
    srcs = ["dao.go"],
    importpath = "go-common/app/interface/main/app-intl/dao/region",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-intl/conf",
        "//app/interface/main/app-intl/model/region",
        "//library/database/sql",
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
