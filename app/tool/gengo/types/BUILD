package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "types_test",
    srcs = [
        "comments_test.go",
        "flatten_test.go",
        "types_test.go",
    ],
    embed = [":types"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "types",
    srcs = [
        "comments.go",
        "doc.go",
        "flatten.go",
        "types.go",
    ],
    importpath = "go-common/app/tool/gengo/types",
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
