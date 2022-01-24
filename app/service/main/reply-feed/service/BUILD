package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "admin.go",
        "reply.go",
        "service.go",
    ],
    importpath = "go-common/app/service/main/reply-feed/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/reply-feed/api",
        "//app/service/main/reply-feed/conf",
        "//app/service/main/reply-feed/dao",
        "//app/service/main/reply-feed/model",
        "//library/log",
        "//library/net/netutil",
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
