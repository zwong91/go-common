package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "cache.go",
        "databus.go",
        "handler.go",
        "handler_resource.go",
        "handler_task.go",
        "monitor.go",
        "report.go",
        "service.go",
        "task_job.go",
        "task_log.go",
        "weight.go",
    ],
    importpath = "go-common/app/job/main/aegis/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/aegis/conf",
        "//app/job/main/aegis/dao",
        "//app/job/main/aegis/dao/email",
        "//app/job/main/aegis/dao/monitor",
        "//app/job/main/aegis/model",
        "//app/job/main/aegis/model/monitor",
        "//app/service/main/account/api",
        "//app/service/main/up/api/v1:api",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//library/queue/databus/databusutil",
        "//library/queue/databus/report",
        "//library/xstr",
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

go_test(
    name = "service_test",
    srcs = [
        "monitor_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/aegis/conf",
        "//app/job/main/aegis/model/monitor",
        "//app/service/main/account/api",
        "//library/log",
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
