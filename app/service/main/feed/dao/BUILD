package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "bangumi_test.go",
        "dao_test.go",
        "memcache_test.go",
        "redis_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//app/service/main/feed/conf",
        "//app/service/main/feed/model",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "bangumi.go",
        "dao.go",
        "memcache.go",
        "redis.go",
    ],
    importpath = "go-common/app/service/main/feed/dao",
    tags = ["automanaged"],
    deps = [
        "//app/interface/openplatform/article/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//app/service/main/feed/conf",
        "//app/service/main/feed/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
        "//library/sync/errgroup",
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
