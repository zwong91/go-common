package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "binlog.go",
        "binlog_elec.go",
        "elec_syncdb.go",
        "limited_list.go",
        "repaire.go",
        "service.go",
        "task.go",
        "task_account_biz.go",
        "task_account_user.go",
        "task_bill_daily.go",
        "task_bill_monthly.go",
        "task_shell_recharge.go",
        "tool.go",
    ],
    importpath = "go-common/app/job/main/ugcpay/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/ugcpay/conf",
        "//app/job/main/ugcpay/dao",
        "//app/job/main/ugcpay/model",
        "//app/job/main/ugcpay/service/pay",
        "//library/database/sql",
        "//library/log",
        "//library/queue/databus",
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
    srcs = [
        ":package-srcs",
        "//app/job/main/ugcpay/service/pay:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "service_test",
    srcs = ["service_test.go"],
    embed = [":service"],
    tags = ["automanaged"],
)
