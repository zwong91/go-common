package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "vegas_test",
    srcs = ["vegas_test.go"],
    embed = [":vegas"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//library/rate"],
)

go_library(
    name = "vegas",
    srcs = [
        "sample.go",
        "vegas.go",
    ],
    importpath = "go-common/library/rate/vegas",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/rate"],
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
