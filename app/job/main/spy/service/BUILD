package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "block_test.go",
        "rebuild_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/spy/conf",
        "//app/job/main/spy/model",
        "//library/cache/redis",
        "//library/log",
        "//vendor/github.com/robfig/cron",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "block.go",
        "rebuild.go",
        "report.go",
        "service.go",
        "stat.go",
    ],
    importpath = "go-common/app/job/main/spy/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/spy/conf",
        "//app/job/main/spy/dao",
        "//app/job/main/spy/model",
        "//app/service/main/spy/model",
        "//app/service/main/spy/rpc/client",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//library/stat/prom",
        "//library/time",
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
