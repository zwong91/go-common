package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "extra_func_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/videoup/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "audit.go",
        "databus.go",
        "delay.go",
        "extra_func.go",
        "history.go",
        "result.go",
        "retry.go",
        "round.go",
        "send.go",
        "service.go",
        "shot.go",
        "stat.go",
        "video.go",
    ],
    importpath = "go-common/app/job/main/videoup/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/videoup/conf",
        "//app/job/main/videoup/dao/activity",
        "//app/job/main/videoup/dao/archive",
        "//app/job/main/videoup/dao/bvc",
        "//app/job/main/videoup/dao/manager",
        "//app/job/main/videoup/dao/message",
        "//app/job/main/videoup/dao/monitor",
        "//app/job/main/videoup/dao/redis",
        "//app/job/main/videoup/model/archive",
        "//app/job/main/videoup/model/manager",
        "//app/job/main/videoup/model/message",
        "//app/job/main/videoup/model/redis",
        "//app/service/main/account/api",
        "//library/conf/env",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//library/stat/prom",
        "//library/time",
        "//vendor/github.com/pkg/errors",
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
