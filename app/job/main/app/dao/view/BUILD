package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "view_test",
    srcs = [
        "dao_test.go",
        "memcache_test.go",
        "redis_test.go",
    ],
    embed = [":view"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/app/conf",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "view",
    srcs = [
        "dao.go",
        "memcache.go",
        "redis.go",
    ],
    importpath = "go-common/app/job/main/app/dao/view",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/app/conf",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/net/http/blademaster",
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
