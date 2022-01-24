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
        "memcache_test.go",
        "mysql_test.go",
        "redis_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/spy/conf",
        "//app/job/main/spy/model",
        "//library/database/sql",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "account.go",
        "dao.go",
        "hbase.go",
        "memcache.go",
        "mysql.go",
        "redis.go",
    ],
    importpath = "go-common/app/job/main/spy/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/spy/conf",
        "//app/job/main/spy/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/hbase.v2:hbase_v2",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/github.com/pkg/errors",
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
