package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "client_test.go",
        "dao_test.go",
        "databus_test.go",
        "memcache_test.go",
        "mysql_test.go",
        "push_test.go",
        "redis_test.go",
        "report_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/push/conf",
        "//app/service/main/push/dao/oppo",
        "//app/service/main/push/model",
        "//library/cache/redis",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "client.go",
        "dao.go",
        "databus.go",
        "mc.cache.go",
        "memcache.go",
        "mysql.go",
        "mysql_report.go",
        "mysql_setting.go",
        "push.go",
        "redis.go",
        "report.go",
    ],
    importpath = "go-common/app/service/main/push/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/push/conf",
        "//app/service/main/push/dao/apns2",
        "//app/service/main/push/dao/fcm",
        "//app/service/main/push/dao/huawei",
        "//app/service/main/push/dao/jpush",
        "//app/service/main/push/dao/mi",
        "//app/service/main/push/dao/oppo",
        "//app/service/main/push/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/conf/env",
        "//library/database/sql",
        "//library/log",
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
    srcs = [
        ":package-srcs",
        "//app/service/main/push/dao/apns2:all-srcs",
        "//app/service/main/push/dao/fcm:all-srcs",
        "//app/service/main/push/dao/huawei:all-srcs",
        "//app/service/main/push/dao/jpush:all-srcs",
        "//app/service/main/push/dao/mi:all-srcs",
        "//app/service/main/push/dao/oppo:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
