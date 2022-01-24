package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "recdata.go",
    ],
    importpath = "go-common/app/service/live/xroom-feed/internal/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/xroom-feed/internal/model",
        "//library/cache/memcache",
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
