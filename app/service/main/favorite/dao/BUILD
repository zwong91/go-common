package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "databus_test.go",
        "memcache_test.go",
        "mysql_test.go",
        "redis_test.go",
        "search_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/favorite/conf",
        "//app/service/main/favorite/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "databus.go",
        "memcache.go",
        "mysql.go",
        "redis.go",
        "search.go",
    ],
    importpath = "go-common/app/service/main/favorite/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/favorite/conf",
        "//app/service/main/favorite/model",
        "//app/service/main/riot-search/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/queue/databus",
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
