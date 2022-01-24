package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "reply_test",
    srcs = [
        "admin_test.go",
        "business_test.go",
        "content_test.go",
        "dao_test.go",
        "databus_test.go",
        "memcache_test.go",
        "redis_test.go",
        "reply_test.go",
        "report_test.go",
        "subject_test.go",
    ],
    embed = [":reply"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/reply/conf",
        "//app/job/main/reply/model/reply",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "reply",
    srcs = [
        "admin.go",
        "business.go",
        "content.go",
        "dao.go",
        "databus.go",
        "fold.go",
        "memcache.go",
        "redis.go",
        "reply.go",
        "report.go",
        "subject.go",
    ],
    importpath = "go-common/app/job/main/reply/dao/reply",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/reply/conf",
        "//app/job/main/reply/model/reply",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/queue/databus",
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
