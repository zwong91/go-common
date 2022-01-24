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
        "hbase_test.go",
        "mysql_test.go",
        "redis_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/history/model",
        "//app/job/main/history/conf",
        "//app/service/main/history/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "hbase.go",
        "mysql.go",
        "redis.go",
    ],
    importpath = "go-common/app/job/main/history/dao",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/history/model",
        "//app/job/main/history/conf",
        "//app/service/main/history/model",
        "//library/cache/redis",
        "//library/database/hbase.v2:hbase_v2",
        "//library/database/tidb",
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
