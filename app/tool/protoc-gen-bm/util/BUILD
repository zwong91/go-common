package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "util_test",
    srcs = ["style_test.go"],
    embed = [":util"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "util",
    srcs = [
        "param.go",
        "style.go",
    ],
    importpath = "go-common/app/tool/protoc-gen-bm/util",
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
