package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "archive_test.go",
        "auth_role_test.go",
        "consumer_test.go",
        "dao_test.go",
        "manager_test.go",
        "oper_test.go",
        "report_test.go",
        "search_test.go",
        "task_dispatch_test.go",
        "task_history_test.go",
        "task_oper_history_test.go",
        "task_qa_test.go",
        "task_qa_video_test.go",
        "task_review_test.go",
        "task_weight_test.go",
        "type_test.go",
        "up_group_test.go",
        "user_test.go",
        "video_test.go",
        "weight_log_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/videoup-task/conf",
        "//app/admin/main/videoup-task/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "archive.go",
        "auth_role.go",
        "consumer.go",
        "dao.go",
        "manager.go",
        "oper.go",
        "report.go",
        "search.go",
        "task_dispatch.go",
        "task_history.go",
        "task_oper_history.go",
        "task_qa.go",
        "task_qa_video.go",
        "task_review.go",
        "task_weight.go",
        "type.go",
        "up_group.go",
        "user.go",
        "video.go",
        "weight_log.go",
    ],
    importpath = "go-common/app/admin/main/videoup-task/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/videoup-task/conf",
        "//app/admin/main/videoup-task/model",
        "//app/admin/main/videoup/model/archive",
        "//app/service/main/account/api",
        "//app/service/main/account/model",
        "//library/cache/redis",
        "//library/database/hbase.v2:hbase_v2",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
        "//library/xstr",
        "//vendor/github.com/tsuna/gohbase/hrpc",
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
