package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "deliver_test",
    srcs = ["deliver_test.go"],
    embed = [":deliver"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "deliver",
    srcs = ["deliver.go"],
    importpath = "go-common/app/service/main/dapper/pkg/deliver",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/log"],
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
