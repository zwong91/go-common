package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "archive_test.go",
        "excitation_test.go",
        "monitor_test.go",
        "reply_test.go",
        "retry_test.go",
        "service_test.go",
        "tag_test.go",
        "ups_test.go",
        "video_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/videoup-report/conf",
        "//app/job/main/videoup-report/model/archive",
        "//app/job/main/videoup-report/model/email",
        "//app/job/main/videoup-report/model/manager",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "archive.go",
        "archive_result.go",
        "databus.go",
        "email.go",
        "excitation.go",
        "extra_func.go",
        "flow.go",
        "monitor.go",
        "reply.go",
        "retry.go",
        "service.go",
        "tag.go",
        "task.go",
        "task_assign.go",
        "task_dispatch.go",
        "task_weight.go",
        "track.go",
        "ups.go",
        "video.go",
    ],
    importpath = "go-common/app/job/main/videoup-report/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/videoup-report/conf",
        "//app/job/main/videoup-report/dao/archive",
        "//app/job/main/videoup-report/dao/data",
        "//app/job/main/videoup-report/dao/email",
        "//app/job/main/videoup-report/dao/hbase",
        "//app/job/main/videoup-report/dao/manager",
        "//app/job/main/videoup-report/dao/mission",
        "//app/job/main/videoup-report/dao/redis",
        "//app/job/main/videoup-report/dao/tag",
        "//app/job/main/videoup-report/model/archive",
        "//app/job/main/videoup-report/model/email",
        "//app/job/main/videoup-report/model/manager",
        "//app/job/main/videoup-report/model/monitor",
        "//app/job/main/videoup-report/model/task",
        "//app/job/main/videoup-report/model/utils",
        "//app/service/main/account/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/up/api/v1:api",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/queue/databus",
        "//library/queue/databus/report",
        "//library/sync/errgroup",
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
