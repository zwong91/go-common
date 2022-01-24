package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "expr_test",
    srcs = ["expr_test.go"],
    embed = [":expr"],
    tags = ["automanaged"],
)

go_library(
    name = "expr",
    srcs = [
        "ast.go",
        "check.go",
        "eval.go",
        "expr.go",
    ],
    importpath = "go-common/app/service/live/zeus/expr",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
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
