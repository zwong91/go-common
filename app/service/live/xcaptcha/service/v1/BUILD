package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = ["xCaptcha.go"],
    importpath = "go-common/app/service/live/xcaptcha/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/rtc/common",
        "//app/service/live/xcaptcha/api/grpc/v1:grpc",
        "//app/service/live/xcaptcha/conf",
        "//app/service/live/xcaptcha/dao",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
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
