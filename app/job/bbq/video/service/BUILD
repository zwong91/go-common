package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "binlogv.go",
        "binlogvr.go",
        "bvcsub.go",
        "cmsvideo.go",
        "comment.go",
        "email.go",
        "ftp.go",
        "notice.go",
        "service.go",
        "user.go",
        "video.go",
    ],
    importpath = "go-common/app/job/bbq/video/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/bbq/video/conf",
        "//app/job/bbq/video/dao",
        "//app/job/bbq/video/model",
        "//app/service/bbq/common",
        "//app/service/bbq/notice-service/api/v1:api",
        "//app/service/bbq/search/api/grpc/v1:grpc",
        "//app/service/bbq/topic/api",
        "//app/service/bbq/video/api/grpc/v1:grpc",
        "//library/conf/env",
        "//library/log",
        "//library/queue/databus",
        "//library/sync/errgroup",
        "//vendor/github.com/ftp-master",
        "//vendor/github.com/robfig/cron",
        "//vendor/gopkg.in/gomail.v2:gomail_v2",
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
