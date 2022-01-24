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
        "mysql.go",
        "redis.go",
        "redisKey.go",
    ],
    importpath = "go-common/app/service/live/dao-anchor/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/common/live/library/lrucache",
        "//app/service/live/dao-anchor/api/grpc/v1:grpc",
        "//app/service/live/dao-anchor/conf",
        "//app/service/live/dao-anchor/model",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//library/sync/errgroup",
        "//vendor/github.com/json-iterator/go",
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
        "dao_test.go",
        "mysql_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/dao-anchor/api/grpc/v1:grpc",
        "//app/service/live/dao-anchor/conf",
        "//app/service/live/dao-anchor/model",
        "//library/log",
        "//library/queue/databus",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
