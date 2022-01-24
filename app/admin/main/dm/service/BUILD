package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "advance_test.go",
        "dm_test.go",
        "filter_test.go",
        "mask_test.go",
        "monitor_test.go",
        "report_test.go",
        "service_test.go",
        "subject_test.go",
        "task_test.go",
        "transfer_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/dm/conf",
        "//app/admin/main/dm/model",
        "//app/admin/main/dm/model/oplog",
        "//library/log",
        "//library/queue/databus/report",
        "//vendor/github.com/davecgh/go-spew/spew",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "advance.go",
        "dm.go",
        "filter.go",
        "mask.go",
        "monitor.go",
        "report.go",
        "service.go",
        "shield.go",
        "subject.go",
        "subtitle.go",
        "subtitle_status.go",
        "subtitle_subject.go",
        "task.go",
        "transfer.go",
    ],
    importpath = "go-common/app/admin/main/dm/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/dm/conf",
        "//app/admin/main/dm/dao",
        "//app/admin/main/dm/dao/oplog",
        "//app/admin/main/dm/model",
        "//app/admin/main/dm/model/oplog",
        "//app/service/main/account/api",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//library/queue/databus/report",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
        "//library/xstr",
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
