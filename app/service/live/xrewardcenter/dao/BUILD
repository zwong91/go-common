package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = ["dao.go"],
    importpath = "go-common/app/service/live/xrewardcenter/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/gift/api/liverpc",
        "//app/service/live/room/api/liverpc",
        "//app/service/live/xrewardcenter/conf",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/net/rpc/liverpc",
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
        "//app/service/live/xrewardcenter/dao/anchorReward:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
