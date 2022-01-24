package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "batchwrite_test",
    srcs = ["batchwrite_test.go"],
    embed = [":batchwrite"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//app/service/main/dapper/model"],
)

go_library(
    name = "batchwrite",
    srcs = ["batchwrite.go"],
    importpath = "go-common/app/service/main/dapper/pkg/batchwrite",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/dapper/model",
        "//library/log",
        "//vendor/github.com/pkg/errors",
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
