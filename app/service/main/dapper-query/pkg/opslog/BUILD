package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "opslog_test",
    srcs = ["opslog_test.go"],
    embed = [":opslog"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "opslog",
    srcs = ["opslog.go"],
    importpath = "go-common/app/service/main/dapper-query/pkg/opslog",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/pkg/errors"],
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
