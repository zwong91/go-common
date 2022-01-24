package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "associate_test.go",
        "autorenew_test.go",
        "coupon_test.go",
        "frozen_test.go",
        "order_test.go",
        "salary_test.go",
        "service_test.go",
        "sync_test.go",
        "vip_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/vip/conf",
        "//app/job/main/vip/model",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "associate.go",
        "autorenew.go",
        "bcoin.go",
        "check_data.go",
        "coupon.go",
        "frozen.go",
        "message.go",
        "order.go",
        "push.go",
        "salary.go",
        "salary_data.go",
        "service.go",
        "sync.go",
        "vip.go",
    ],
    importpath = "go-common/app/job/main/vip/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/vip/conf",
        "//app/job/main/vip/dao",
        "//app/job/main/vip/model",
        "//app/service/main/coupon/model",
        "//app/service/main/coupon/rpc/client",
        "//app/service/main/vip/api",
        "//app/service/main/vip/model",
        "//library/database/sql",
        "//library/log",
        "//library/queue/databus",
        "//library/sync/errgroup",
        "//library/time",
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
