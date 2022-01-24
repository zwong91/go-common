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
        "dao_test.go",
        "hbase_test.go",
        "mysql_test.go",
        "redis_test.go",
        "vip_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/figure/conf",
        "//app/job/main/figure/model",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "account.go",
        "blocked.go",
        "danmaku.go",
        "dao.go",
        "hbase.go",
        "mysql.go",
        "redis.go",
        "vip.go",
    ],
    importpath = "go-common/app/job/main/figure/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/figure/conf",
        "//app/job/main/figure/model",
        "//library/cache/redis",
        "//library/database/hbase.v2:hbase_v2",
        "//library/database/sql",
        "//library/log",
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
