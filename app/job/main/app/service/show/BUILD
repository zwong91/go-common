package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "show_test",
    srcs = ["show_test.go"],
    embed = [":show"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/app/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "show",
    srcs = ["show.go"],
    importpath = "go-common/app/job/main/app/service/show",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/app/conf",
        "//app/job/main/app/dao/show",
        "//app/job/main/app/model",
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
