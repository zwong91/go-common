package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "now_test",
    srcs = ["now_test.go"],
    embed = [":now"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "now",
    srcs = [
        "main.go",
        "now.go",
    ],
    importpath = "go-common/app/admin/main/up/util/now",
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
