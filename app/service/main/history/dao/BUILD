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
        "redis_test.go",
        "tidb_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/history/api/grpc",
        "//app/service/main/history/conf",
        "//app/service/main/history/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "databus.go",
        "redis.go",
        "tidb.go",
    ],
    importpath = "go-common/app/service/main/history/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/history/api/grpc",
        "//app/service/main/history/conf",
        "//app/service/main/history/model",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/database/tidb",
        "//library/log",
        "//library/queue/databus",
        "//library/stat/prom",
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
