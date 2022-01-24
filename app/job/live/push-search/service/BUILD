package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "attention_notify.go",
        "common.go",
        "roominfo_notify.go",
        "service.go",
        "uname_notify.go",
    ],
    importpath = "go-common/app/job/live/push-search/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/live/push-search/conf",
        "//app/job/live/push-search/dao",
        "//app/job/live/push-search/model",
        "//app/service/live/relation/api/liverpc/v1:liverpc",
        "//app/service/live/room/api/liverpc/v1:liverpc",
        "//app/service/live/room/api/liverpc/v2:liverpc",
        "//app/service/live/user/api/liverpc/v3:liverpc",
        "//app/service/main/account/api",
        "//library/log",
        "//library/net/rpc/liverpc/context",
        "//library/queue/databus",
        "//library/sync/errgroup",
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
        "//app/job/live/push-search/service/migrate:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
