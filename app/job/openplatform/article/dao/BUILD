package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "activity_test.go",
        "dao_test.go",
        "game_test.go",
        "mysql_test.go",
        "redis_test.go",
        "reply_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/openplatform/article/model",
        "//app/job/openplatform/article/conf",
        "//library/cache/redis",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "activity.go",
        "art_redis.go",
        "dao.cache.go",
        "dao.go",
        "dynamic.go",
        "flow.go",
        "game.go",
        "media.go",
        "mysql.go",
        "purge.go",
        "recommend.go",
        "redis.go",
        "reply.go",
    ],
    importpath = "go-common/app/job/openplatform/article/dao",
    tags = ["automanaged"],
    deps = [
        "//app/interface/openplatform/article/model",
        "//app/job/openplatform/article/conf",
        "//app/job/openplatform/article/model",
        "//library/cache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/queue/databus",
        "//library/stat/prom",
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
