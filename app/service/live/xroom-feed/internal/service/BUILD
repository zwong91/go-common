package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "pool_conf.go",
        "rec_pool.go",
        "service.go",
    ],
    importpath = "go-common/app/service/live/xroom-feed/internal/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/xroom-feed/api",
        "//app/service/live/xroom-feed/internal/dao",
        "//app/service/live/xroom-feed/internal/model",
        "//library/conf/paladin",
        "//library/log",
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
