package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "cache_update_test.go",
        "dao_test.go",
        "follower_test.go",
        "following_test.go",
        "mysql_test.go",
        "stat_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/relation/conf",
        "//app/job/main/relation/model",
        "//app/service/main/relation/model",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "cache_update.go",
        "dao.go",
        "follower.go",
        "following.go",
        "mysql.go",
        "stat.go",
    ],
    importpath = "go-common/app/job/main/relation/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/relation/conf",
        "//app/job/main/relation/model",
        "//app/job/main/relation/model/i64b",
        "//app/service/main/relation/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/time",
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
