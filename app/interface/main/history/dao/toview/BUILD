package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "toview_test",
    srcs = [
        "dao_test.go",
        "hbase_test.go",
        "redis_test.go",
    ],
    embed = [":toview"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/history/conf",
        "//app/interface/main/history/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "toview",
    srcs = [
        "dao.go",
        "hbase.go",
        "redis.go",
    ],
    importpath = "go-common/app/interface/main/history/dao/toview",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/history/conf",
        "//app/interface/main/history/model",
        "//library/cache/redis",
        "//library/database/hbase.v2:hbase_v2",
        "//library/log",
        "//vendor/github.com/tsuna/gohbase/hrpc",
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
