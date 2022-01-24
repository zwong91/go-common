load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "hbase.go",
        "http.go",
        "memcache.go",
        "redis.go",
    ],
    importpath = "go-common/app/service/main/secure/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/secure/conf",
        "//app/service/main/secure/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/hbase.v2:hbase_v2",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/time",
        "//vendor/github.com/tsuna/gohbase/hrpc",
    ],
)

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "hbase_test.go",
        "http_test.go",
        "memcache_test.go",
        "redis_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/secure/conf",
        "//app/service/main/secure/model",
        "//library/cache/memcache",
        "//library/time",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
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
