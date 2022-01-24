package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "roomAdmin_test",
    srcs = [
        "dao_test.go",
        "mc.cache_test.go",
        "mc_test.go",
        "mysql_test.go",
        "redis_test.go",
    ],
    embed = [":roomAdmin"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/xuser/conf",
        "//app/service/live/xuser/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "roomAdmin",
    srcs = [
        "dao.go",
        "mc.cache.go",
        "mc.go",
        "mysql.go",
        "redis.go",
    ],
    importpath = "go-common/app/service/live/xuser/dao/roomAdmin",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/banned_service/api/liverpc/v1:liverpc",
        "//app/service/live/fans_medal/api/liverpc/v2:liverpc",
        "//app/service/live/room/api/liverpc/v1:liverpc",
        "//app/service/live/room/api/liverpc/v2:liverpc",
        "//app/service/live/xuser/api/grpc/v1:grpc",
        "//app/service/live/xuser/conf",
        "//app/service/live/xuser/dao",
        "//app/service/live/xuser/model",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/orm",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
        "//vendor/github.com/jinzhu/gorm",
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
