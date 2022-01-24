package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "api.go",
        "dao.go",
        "pub.go",
    ],
    importpath = "go-common/app/job/live/push-search/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/live/push-search/conf",
        "//app/service/live/relation/api/liverpc",
        "//app/service/live/room/api/liverpc",
        "//app/service/live/user/api/liverpc",
        "//library/database/hbase.v2:hbase_v2",
        "//library/log",
        "//library/net/rpc/liverpc",
        "//library/queue/databus",
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
    srcs = [
        ":package-srcs",
        "//app/job/live/push-search/dao/migrate:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
