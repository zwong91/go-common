package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "music_test",
    srcs = ["service_test.go"],
    embed = [":music"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/service",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "music",
    srcs = ["service.go"],
    importpath = "go-common/app/interface/main/creative/service/music",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/account",
        "//app/interface/main/creative/dao/archive",
        "//app/interface/main/creative/dao/material",
        "//app/interface/main/creative/dao/music",
        "//app/interface/main/creative/dao/up",
        "//app/interface/main/creative/model/app",
        "//app/interface/main/creative/model/music",
        "//app/interface/main/creative/model/search",
        "//app/interface/main/creative/service",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/queue/databus/report",
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
