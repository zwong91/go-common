package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "collect",
    srcs = ["collect.go"],
    importpath = "go-common/app/service/main/dapper/pkg/collect",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//app/service/main/dapper/pkg/process"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//app/service/main/dapper/pkg/collect/kafkacollect:all-srcs",
        "//app/service/main/dapper/pkg/collect/tcpcollect:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
