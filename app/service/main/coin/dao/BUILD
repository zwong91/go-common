package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "cache_test.go",
        "dao.cache_test.go",
        "dao_test.go",
        "databus_test.go",
        "log_test.go",
        "mc.cache_test.go",
        "mysql_coin_test.go",
        "mysql_test.go",
        "redis_test.go",
        "relation_test.go",
        "tag_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/coin/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "cache.go",
        "dao.cache.go",
        "dao.go",
        "databus.go",
        "log.go",
        "mc.cache.go",
        "mysql.go",
        "mysql_coin.go",
        "redis.go",
        "relation.go",
        "tag.go",
    ],
    importpath = "go-common/app/service/main/coin/dao",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/coin/api",
        "//app/service/main/coin/conf",
        "//app/service/main/coin/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/elastic",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/queue/databus",
        "//library/queue/databus/report",
        "//library/stat/prom",
        "//library/sync/pipeline/fanout",
        "//vendor/github.com/pkg/errors",
        "//vendor/golang.org/x/sync/singleflight",
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
