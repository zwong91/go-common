package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "appeal_test.go",
        "blocked_mc_test.go",
        "blocked_redis_test.go",
        "blocked_test.go",
        "case_mc_test.go",
        "case_redis_test.go",
        "case_test.go",
        "dao_test.go",
        "databus_test.go",
        "http_test.go",
        "jury_test.go",
        "manager_test.go",
        "mc_test.go",
        "memcache_test.go",
        "mysql_test.go",
        "opinion_test.go",
        "redis_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/credit/conf",
        "//app/interface/main/credit/model",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "appeal.go",
        "blocked.go",
        "blocked_mc.go",
        "blocked_redis.go",
        "case.go",
        "case_mc.go",
        "case_redis.go",
        "dao.go",
        "databus.go",
        "http.go",
        "jury.go",
        "manager.go",
        "memcache.go",
        "mysql.go",
        "opinion.go",
        "redis.go",
    ],
    importpath = "go-common/app/interface/main/credit/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/credit/conf",
        "//app/interface/main/credit/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/queue/databus",
        "//library/sync/errgroup",
        "//library/time",
        "//library/xstr",
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
