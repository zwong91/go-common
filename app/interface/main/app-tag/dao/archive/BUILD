package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "archive_test",
    srcs = ["dao_test.go"],
    embed = [":archive"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-tag/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "archive",
    srcs = ["dao.go"],
    importpath = "go-common/app/interface/main/app-tag/dao/archive",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-tag/conf",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
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
