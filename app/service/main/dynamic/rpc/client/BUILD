package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "client",
    srcs = ["dynamic.go"],
    importpath = "go-common/app/service/main/dynamic/rpc/client",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/api",
        "//app/service/main/dynamic/model",
        "//library/net/rpc",
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
