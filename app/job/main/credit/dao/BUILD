package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "auto_case_test.go",
        "blocked_test.go",
        "case_mc_test.go",
        "case_test.go",
        "dao_test.go",
        "depend_test.go",
        "jury_test.go",
        "memcache_test.go",
        "mysql_test.go",
        "redis_test.go",
        "reply_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/credit/conf",
        "//app/job/main/credit/model",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "auto_case.go",
        "blocked.go",
        "case.go",
        "case_mc.go",
        "dao.go",
        "depend.go",
        "jury.go",
        "memcache.go",
        "mysql.go",
        "redis.go",
        "reply.go",
    ],
    importpath = "go-common/app/job/main/credit/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/credit/conf",
        "//app/job/main/credit/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
