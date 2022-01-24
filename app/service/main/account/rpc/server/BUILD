package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "server",
    srcs = [
        "rpc.go",
        "rpc3.go",
    ],
    importpath = "go-common/app/service/main/account/rpc/server",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/account/api",
        "//app/service/main/account/conf",
        "//app/service/main/account/model",
        "//app/service/main/account/service",
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
