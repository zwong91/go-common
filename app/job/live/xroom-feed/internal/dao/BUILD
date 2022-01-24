package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "cache.go",
        "conf.go",
        "dao.go",
        "rediskey.go",
    ],
    importpath = "go-common/app/job/live/xroom-feed/internal/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/live/xroom-feed/internal/model",
        "//app/service/live/dao-anchor/api/grpc/v1:grpc",
        "//library/cache/redis",
        "//library/conf/paladin",
        "//library/database/sql",
        "//library/log",
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
