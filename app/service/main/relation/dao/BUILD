package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "audit.go",
        "dao.go",
        "localcache.go",
        "memcache.go",
        "mysql.go",
        "prompt_redis.go",
        "recent_redis.go",
        "redis.go",
        "relation_log.go",
    ],
    importpath = "go-common/app/service/main/relation/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/relation/conf",
        "//app/service/main/relation/model",
        "//app/service/main/relation/model/i64b",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/queue/databus/report",
        "//library/stat/prom",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/bluele/gcache",
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

go_test(
    name = "dao_test",
    srcs = [
        "audit_test.go",
        "dao_test.go",
        "localcache_test.go",
        "memcache_test.go",
        "mysql_test.go",
        "prompt_redis_test.go",
        "recent_redis_test.go",
        "redis_test.go",
        "relation_log_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/relation/conf",
        "//app/service/main/relation/model",
        "//app/service/main/relation/model/i64b",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
