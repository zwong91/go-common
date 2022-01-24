package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "backup-stream.go",
        "cache.go",
        "dao.cache.go",
        "dao.go",
        "http.go",
        "localcache.go",
        "main-stream.go",
        "mysql.go",
        "notify.go",
        "official-stream.go",
        "redis.go",
        "stream-change-log.go",
        "upstream-summary.go",
    ],
    importpath = "go-common/app/service/video/stream-mng/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/video/stream-mng/common",
        "//app/service/video/stream-mng/conf",
        "//app/service/video/stream-mng/model",
        "//library/cache",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/conf/env",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/stat/prom",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
        "//vendor/github.com/bluele/gcache",
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

go_test(
    name = "dao_test",
    srcs = [
        "backup-stream_test.go",
        "cache_test.go",
        "dao.cache_test.go",
        "dao_test.go",
        "http_test.go",
        "localcache_test.go",
        "main-stream_test.go",
        "mysql_test.go",
        "notify_test.go",
        "official-stream_test.go",
        "redis_test.go",
        "stream-change-log_test.go",
        "upstream-summary_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/video/stream-mng/conf",
        "//app/service/video/stream-mng/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
