package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "language_test",
    srcs = ["language_test.go"],
    embed = [":language"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/app/conf",
        "//app/admin/main/app/model/language",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "language",
    srcs = ["language.go"],
    importpath = "go-common/app/admin/main/app/service/language",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/app/conf",
        "//app/admin/main/app/dao/language",
        "//app/admin/main/app/model/language",
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
