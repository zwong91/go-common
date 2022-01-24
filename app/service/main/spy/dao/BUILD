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
        "audit_test.go",
        "bilibili_card_test.go",
        "dao_test.go",
        "databus_test.go",
        "memcache_test.go",
        "mysql_test.go",
        "qcloud_test.go",
        "redis_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/spy/conf",
        "//app/service/main/spy/model",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "account.go",
        "audit.go",
        "bilibili_card.go",
        "dao.go",
        "databus.go",
        "memcache.go",
        "mysql.go",
        "qcloud.go",
        "redis.go",
    ],
    importpath = "go-common/app/service/main/spy/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/spy/conf",
        "//app/service/main/spy/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/queue/databus",
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
