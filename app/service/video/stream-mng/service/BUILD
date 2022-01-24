package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "backup.go",
        "change-log.go",
        "change-src.go",
        "check-live.go",
        "clear-cache.go",
        "clear-retweet.go",
        "common.go",
        "main-stream.go",
        "notify.go",
        "official.go",
        "screen-shot.go",
        "service.go",
        "stream.go",
        "stream-cut.go",
        "stream-info.go",
        "stream-mask.go",
        "stream-validate.go",
        "upstream-rtmp.go",
        "upstream-summary.go",
    ],
    importpath = "go-common/app/service/video/stream-mng/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/location/api",
        "//app/service/video/stream-mng/common",
        "//app/service/video/stream-mng/conf",
        "//app/service/video/stream-mng/dao",
        "//app/service/video/stream-mng/model",
        "//library/conf/env",
        "//library/log",
        "//library/net/metadata",
        "//library/net/rpc/warden",
        "//library/sync/pipeline/fanout",
        "//library/time",
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
