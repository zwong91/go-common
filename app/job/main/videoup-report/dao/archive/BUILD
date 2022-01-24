package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "archive_test",
    srcs = [
        "addit_test.go",
        "archive_oper_test.go",
        "archive_recheck_test.go",
        "archive_test.go",
        "audit_test.go",
        "config_test.go",
        "dao_test.go",
        "flow_design_test.go",
        "task_test.go",
        "type_test.go",
        "video_test.go",
        "video_track_test.go",
    ],
    embed = [":archive"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/videoup-report/conf",
        "//app/job/main/videoup-report/model/archive",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "archive",
    srcs = [
        "addit.go",
        "archive.go",
        "archive_delay.go",
        "archive_oper.go",
        "archive_recheck.go",
        "audit.go",
        "config.go",
        "dao.go",
        "flow_design.go",
        "report.go",
        "task.go",
        "task_consumer.go",
        "task_dispatch.go",
        "task_dispatch_done.go",
        "task_dispatch_extend.go",
        "task_json_config.go",
        "task_oper_history.go",
        "track.go",
        "type.go",
        "video.go",
        "video_track.go",
    ],
    importpath = "go-common/app/job/main/videoup-report/dao/archive",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/videoup-report/conf",
        "//app/job/main/videoup-report/model/archive",
        "//app/job/main/videoup-report/model/task",
        "//app/job/main/videoup-report/model/utils",
        "//library/database/sql",
        "//library/log",
        "//library/xstr",
        "//vendor/github.com/dgryski/go-farm",
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
