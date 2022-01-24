package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "history_test",
    srcs = [
        "dao_test.go",
        "databus_test.go",
        "exp_test.go",
        "grpc_test.go",
        "hbase_test.go",
        "redis_test.go",
    ],
    embed = [":history"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/history/conf",
        "//app/interface/main/history/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "history",
    srcs = [
        "dao.go",
        "databus.go",
        "exp.go",
        "grpc.go",
        "hbase.go",
        "redis.go",
    ],
    importpath = "go-common/app/interface/main/history/dao/history",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/history/conf",
        "//app/interface/main/history/model",
        "//app/service/openplatform/pgc-season/api/grpc/episode/v1:episode",
        "//library/cache",
        "//library/cache/redis",
        "//library/database/hbase.v2:hbase_v2",
        "//library/log",
        "//library/net/metadata",
        "//library/queue/databus",
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
