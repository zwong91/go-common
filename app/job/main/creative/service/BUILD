package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "service_test.go",
        "weeklyhonor_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/model/weeklyhonor",
        "//app/job/main/creative/conf",
        "//app/service/main/up/api/v1:api",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "academy.go",
        "databus.go",
        "service.go",
        "task.go",
        "weeklyhonor.go",
    ],
    importpath = "go-common/app/job/main/creative/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/model/weeklyhonor",
        "//app/interface/openplatform/article/model",
        "//app/job/main/creative/conf",
        "//app/job/main/creative/dao/academy",
        "//app/job/main/creative/dao/archive",
        "//app/job/main/creative/dao/monitor",
        "//app/job/main/creative/dao/newcomer",
        "//app/job/main/creative/dao/weeklyhonor",
        "//app/job/main/creative/model",
        "//app/service/main/archive/api",
        "//app/service/main/up/api/v1:api",
        "//library/conf/env",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//library/sync/errgroup",
        "//library/xstr",
        "//vendor/github.com/robfig/cron",
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
