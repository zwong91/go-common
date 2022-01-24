package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "recommend_test",
    srcs = [
        "api_test.go",
        "dao_test.go",
        "databus_test.go",
        "memcache_test.go",
        "redis_test.go",
    ],
    embed = [":recommend"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-card/model",
        "//app/interface/main/app-card/model/card/ai",
        "//app/interface/main/app-feed/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "recommend",
    srcs = [
        "api.go",
        "dao.go",
        "databus.go",
        "memcache.go",
        "redis.go",
    ],
    importpath = "go-common/app/interface/main/app-feed/dao/recommend",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-card/model",
        "//app/interface/main/app-card/model/card/ai",
        "//app/interface/main/app-feed/conf",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/queue/databus",
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
