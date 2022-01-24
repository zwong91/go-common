package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "jpush",
    srcs = [
        "audience.go",
        "client.go",
        "message.go",
        "notification.go",
        "option.go",
        "payload.go",
    ],
    importpath = "go-common/app/service/bbq/push/dao/jpush",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/stat",
        "//library/stat/prom",
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
