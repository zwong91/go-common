package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "condrule.go",
        "is_black.go",
        "onlinelist.go",
        "reccache.go",
        "rediskey.go",
        "rule.go",
        "service.go",
        "sort.go",
    ],
    importpath = "go-common/app/job/live/xroom-feed/internal/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/live/xroom-feed/internal/dao",
        "//app/job/live/xroom-feed/internal/model",
        "//app/service/live/dao-anchor/api/grpc/v1:grpc",
        "//app/service/live/room/api/liverpc",
        "//app/service/live/room/api/liverpc/v1:liverpc",
        "//library/conf/env",
        "//library/conf/paladin",
        "//library/log",
        "//library/net/metadata",
        "//library/net/rpc/liverpc",
        "//library/net/rpc/warden",
        "//library/sync/errgroup",
        "//library/time",
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
