package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "cltclient",
    srcs = ["clt.go"],
    importpath = "go-common/app/service/main/dapper-query/pkg/cltclient",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/sync/errgroup"],
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
