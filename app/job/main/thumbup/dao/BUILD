package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "databus.go",
        "item_likes_redis.go",
        "memcached.go",
        "redis.go",
        "tidb.go",
        "user_likes_redis.go",
    ],
    importpath = "go-common/app/job/main/thumbup/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/thumbup/conf",
        "//app/job/main/thumbup/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/tidb",
        "//library/log",
        "//library/queue/databus",
        "//library/time",
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

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "databus_test.go",
        "item_likes_redis_test.go",
        "memcached_test.go",
        "redis_test.go",
        "tidb_test.go",
        "user_likes_redis_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/thumbup/conf",
        "//app/job/main/thumbup/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
