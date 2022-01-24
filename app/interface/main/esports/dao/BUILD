package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "cache.go",
        "dao.cache.go",
        "dao.go",
        "mysql.go",
        "pointdata.go",
        "redis.go",
        "search.go",
    ],
    importpath = "go-common/app/interface/main/esports/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/esports/conf",
        "//app/interface/main/esports/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//library/cache/redis",
        "//library/database/elastic",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/stat/prom",
        "//library/sync/pipeline/fanout",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "dao_test",
    srcs = [
        "cache_test.go",
        "dao.cache_test.go",
        "dao_test.go",
        "mysql_test.go",
        "pointdata_test.go",
        "redis_test.go",
        "search_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/esports/conf",
        "//app/interface/main/esports/model",
        "//app/service/main/archive/api",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
