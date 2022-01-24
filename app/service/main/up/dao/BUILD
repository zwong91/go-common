package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = ["utils_test.go"],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/smartystreets/goconvey/convey"],
)

go_library(
    name = "dao",
    srcs = ["utils.go"],
    importpath = "go-common/app/service/main/up/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//app/service/main/up/dao/archive:all-srcs",
        "//app/service/main/up/dao/card:all-srcs",
        "//app/service/main/up/dao/data:all-srcs",
        "//app/service/main/up/dao/global:all-srcs",
        "//app/service/main/up/dao/manager:all-srcs",
        "//app/service/main/up/dao/monitor:all-srcs",
        "//app/service/main/up/dao/up:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
