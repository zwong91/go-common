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
        "db.go",
        "http_client.go",
        "mc.go",
        "redis.go",
    ],
    importpath = "go-common/app/job/main/reply-feed/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/reply-feed/conf",
        "//app/job/main/reply-feed/model",
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

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "db_test.go",
        "http_client_test.go",
        "mc_test.go",
        "redis_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/reply-feed/conf",
        "//app/job/main/reply-feed/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
