package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "databus.go",
        "loader.go",
        "reply_set.go",
        "reply_zset.go",
        "service.go",
        "stat.go",
        "statistics.go",
    ],
    importpath = "go-common/app/job/main/reply-feed/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/reply-feed/conf",
        "//app/job/main/reply-feed/dao",
        "//app/job/main/reply-feed/model",
        "//library/cache/redis",
        "//library/ecode",
        "//library/log",
        "//library/net/netutil",
        "//library/queue/databus",
        "//library/sync/pipeline/fanout",
        "//vendor/github.com/ivpusic/grpool",
        "//vendor/github.com/robfig/cron",
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
