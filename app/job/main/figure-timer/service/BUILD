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
        "//app/job/main/figure-timer/conf",
        "//app/job/main/figure-timer/dao/mock_dao",
        "//app/job/main/figure-timer/model",
        "//library/log",
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "calc.go",
        "fix.go",
        "rank.go",
        "service.go",
    ],
    importpath = "go-common/app/job/main/figure-timer/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/figure-timer/conf",
        "//app/job/main/figure-timer/dao",
        "//app/job/main/figure-timer/model",
        "//library/log",
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
