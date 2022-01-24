package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "danmu_test",
    srcs = ["service_test.go"],
    embed = [":danmu"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/service",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "danmu",
    srcs = [
        "danmu.go",
        "service.go",
        "subtitle.go",
    ],
    importpath = "go-common/app/interface/main/creative/service/danmu",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/account",
        "//app/interface/main/creative/dao/archive",
        "//app/interface/main/creative/dao/danmu",
        "//app/interface/main/creative/dao/elec",
        "//app/interface/main/creative/dao/subtitle",
        "//app/interface/main/creative/model/danmu",
        "//app/interface/main/creative/service",
        "//app/interface/main/dm/model",
        "//app/interface/main/dm2/model",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup",
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
