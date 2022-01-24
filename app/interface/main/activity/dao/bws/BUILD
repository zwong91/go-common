package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "bws",
    srcs = [
        "achievement.go",
        "bws.go",
        "cache.go",
        "dao.cache.go",
        "mc.cache.go",
        "point.go",
    ],
    importpath = "go-common/app/interface/main/activity/dao/bws",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/model/bws",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/stat/prom",
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
    name = "bws_test",
    srcs = [
        "achievement_test.go",
        "bws_test.go",
        "point_test.go",
    ],
    embed = [":bws"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/model/bws",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
