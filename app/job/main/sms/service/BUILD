package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = ["service_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/sms/conf",
        "//app/job/main/sms/model",
        "//app/service/main/sms/model",
        "//library/net/trace",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "callback.go",
        "log.go",
        "service.go",
        "sms.go",
    ],
    importpath = "go-common/app/job/main/sms/service",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/sms/conf",
        "//app/job/main/sms/dao",
        "//app/job/main/sms/dao/chuanglan",
        "//app/job/main/sms/dao/mengwang",
        "//app/job/main/sms/model",
        "//app/service/main/sms/api",
        "//app/service/main/sms/model",
        "//library/conf/env",
        "//library/log",
        "//library/queue/databus",
        "//library/queue/databus/report",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
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
