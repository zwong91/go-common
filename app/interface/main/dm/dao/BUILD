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
        "dao_test.go",
        "databus_test.go",
        "figure_test.go",
        "mysql_protect_test.go",
        "mysql_report_test.go",
        "mysql_test.go",
        "notify_test.go",
        "redis_test.go",
        "search_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/dm/conf",
        "//app/interface/main/dm/model",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "archive.go",
        "dao.go",
        "databus.go",
        "figure.go",
        "mysql.go",
        "mysql_protect.go",
        "mysql_report.go",
        "notify.go",
        "redis.go",
        "search.go",
    ],
    importpath = "go-common/app/interface/main/dm/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/dm/conf",
        "//app/interface/main/dm/model",
        "//library/cache/redis",
        "//library/database/elastic",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/queue/databus",
        "//library/stat/prom",
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
