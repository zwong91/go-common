package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "d_test.go",
        "hbase_test.go",
        "mysql_test.go",
        "redis_test.go",
        "statisitcs_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/push-archive/conf",
        "//app/interface/main/push-archive/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "Proportion.go",
        "dao.go",
        "fan_group.go",
        "hbase.go",
        "message.go",
        "mysql.go",
        "push.go",
        "redis.go",
        "statistics.go",
    ],
    importpath = "go-common/app/interface/main/push-archive/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/push-archive/conf",
        "//app/interface/main/push-archive/model",
        "//library/cache/redis",
        "//library/database/hbase.v2:hbase_v2",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
        "//library/sync/errgroup",
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
