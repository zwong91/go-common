package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "account_notify_test.go",
        "dao_test.go",
        "databus_test.go",
        "exp_test.go",
        "http_test.go",
        "member_log_test.go",
        "member_old_test.go",
        "member_test.go",
        "memcache_test.go",
        "realname_test.go",
        "redis_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/member/conf",
        "//app/job/main/member/model",
        "//library/database/sql",
        "//library/queue/databus",
        "//library/time",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "account_notify.go",
        "dao.go",
        "databus.go",
        "exp.go",
        "hbase.go",
        "http.go",
        "member.go",
        "member_log.go",
        "member_old.go",
        "memcache.go",
        "realname.go",
        "redis.go",
    ],
    importpath = "go-common/app/job/main/member/dao",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/member/conf",
        "//app/job/main/member/dao/block",
        "//app/job/main/member/model",
        "//app/service/main/member/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/queue/databus",
        "//library/queue/databus/report",
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
    srcs = [
        ":package-srcs",
        "//app/job/main/member/dao/block:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
