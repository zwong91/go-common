package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "pendant_test",
    srcs = [
        "dao_test.go",
        "mysql_test.go",
        "redis_test.go",
    ],
    embed = [":pendant"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/usersuit/conf",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "pendant",
    srcs = [
        "dao.go",
        "mysql.go",
        "redis.go",
    ],
    importpath = "go-common/app/job/main/usersuit/dao/pendant",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/usersuit/conf",
        "//app/job/main/usersuit/model",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
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
