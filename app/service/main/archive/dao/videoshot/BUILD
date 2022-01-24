package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "videoshot_test",
    srcs = [
        "dao_test.go",
        "mysql_test.go",
        "redis_test.go",
    ],
    embed = [":videoshot"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "videoshot",
    srcs = [
        "dao.go",
        "mysql.go",
        "redis.go",
    ],
    importpath = "go-common/app/service/main/archive/dao/videoshot",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/conf",
        "//app/service/main/archive/model/videoshot",
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
