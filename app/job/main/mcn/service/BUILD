package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "email.go",
        "mcn.go",
        "recommend.go",
        "service.go",
        "summary.go",
        "up.go",
    ],
    importpath = "go-common/app/job/main/mcn/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/mcn/conf",
        "//app/job/main/mcn/dao",
        "//app/job/main/mcn/model",
        "//app/service/main/account/api",
        "//library/log",
        "//library/sync/pipeline/fanout",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/pkg/errors",
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

go_test(
    name = "service_test",
    srcs = [
        "mcn_test.go",
        "recommend_test.go",
        "service_test.go",
        "summary_test.go",
        "up_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/mcn/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
