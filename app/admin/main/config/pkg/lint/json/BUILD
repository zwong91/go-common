package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "json_test",
    srcs = ["jsonlint_test.go"],
    embed = [":json"],
    tags = ["automanaged"],
)

go_library(
    name = "json",
    srcs = ["jsonlint.go"],
    importpath = "go-common/app/admin/main/config/pkg/lint/json",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//app/admin/main/config/pkg/lint"],
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
