package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "business_test.go",
        "channel_test.go",
        "dao_test.go",
        "grpc_test.go",
        "mysql_test.go",
        "redis_test.go",
        "resource_test.go",
        "tag_prid_redis_test.go",
        "tag_resource_mc_test.go",
        "tag_resource_redis_test.go",
        "tag_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/tag/conf",
        "//app/job/main/tag/model",
        "//app/service/main/archive/api",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "business.go",
        "channel.go",
        "dao.go",
        "grpc.go",
        "mysql.go",
        "redis.go",
        "resource.go",
        "tag.go",
        "tag_prid_redis.go",
        "tag_resource_mc.go",
        "tag_resource_redis.go",
    ],
    importpath = "go-common/app/job/main/tag/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/tag/conf",
        "//app/job/main/tag/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//app/service/main/filter/api/grpc/v1:grpc",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
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
