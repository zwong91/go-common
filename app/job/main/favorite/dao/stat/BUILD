package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "stat",
    srcs = [
        "dao.go",
        "memcache.go",
        "mysql.go",
        "redis.go",
    ],
    importpath = "go-common/app/job/main/favorite/dao/stat",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/favorite/conf",
        "//app/service/main/favorite/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
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
    name = "stat_test",
    srcs = [
        "dao_test.go",
        "memcache_test.go",
        "mysql_test.go",
        "redis_test.go",
    ],
    embed = [":stat"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/favorite/conf",
        "//app/service/main/favorite/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
