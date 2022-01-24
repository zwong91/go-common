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
        "item_likes_redis_test.go",
        "memcached_test.go",
        "redis_test.go",
        "tidb_test.go",
        "user_likes_redis_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/thumbup/conf",
        "//app/service/main/thumbup/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "cache.go",
        "dao.cache.go",
        "dao.go",
        "databus.go",
        "item_likes_redis.go",
        "memcached.go",
        "redis.go",
        "tidb.go",
        "user_likes_redis.go",
    ],
    importpath = "go-common/app/service/main/thumbup/dao",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/thumbup/api",
        "//app/service/main/thumbup/conf",
        "//app/service/main/thumbup/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/tidb",
        "//library/log",
        "//library/queue/databus",
        "//library/stat/prom",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
        "//library/time",
        "//library/xstr",
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
