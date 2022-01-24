package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "gorpc",
    srcs = ["rpc.go"],
    importpath = "go-common/app/service/main/thumbup/server/gorpc",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/thumbup/conf",
        "//app/service/main/thumbup/model",
        "//app/service/main/thumbup/service",
        "//library/net/rpc",
        "//library/net/rpc/context",
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
