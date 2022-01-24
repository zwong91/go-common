package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "account_test.go",
        "credit_test.go",
        "dao_test.go",
        "mask_test.go",
        "mc_filter_test.go",
        "mc_subtitle_test.go",
        "message_test.go",
        "mysql_advance_test.go",
        "mysql_filter_test.go",
        "mysql_mask_test.go",
        "mysql_report_test.go",
        "mysql_subtitle_test.go",
        "mysql_task_test.go",
        "mysql_test.go",
        "mysql_transfer_test.go",
        "search_subtitle_test.go",
        "search_test.go",
        "task_test.go",
        "videoup_test.go",
        "workflow_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/dm/conf",
        "//app/admin/main/dm/model",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "account.go",
        "bfs.go",
        "credit.go",
        "dao.go",
        "mask.go",
        "mc_filter.go",
        "mc_subtitle.go",
        "message.go",
        "mysql.go",
        "mysql_advance.go",
        "mysql_filter.go",
        "mysql_mask.go",
        "mysql_report.go",
        "mysql_subtitle.go",
        "mysql_task.go",
        "mysql_transfer.go",
        "search.go",
        "search_subtitle.go",
        "task.go",
        "videoup.go",
        "workflow.go",
    ],
    importpath = "go-common/app/admin/main/dm/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/dm/conf",
        "//app/admin/main/dm/model",
        "//library/cache/memcache",
        "//library/database/bfs",
        "//library/database/elastic",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/queue/databus",
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
    srcs = [
        ":package-srcs",
        "//app/admin/main/dm/dao/oplog:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
