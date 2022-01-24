package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "pipeline_test",
    srcs = ["pipeline_test.go"],
    embed = [":pipeline"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/net/metadata",
        "//library/time",
    ],
)

go_library(
    name = "pipeline",
    srcs = ["pipeline.go"],
    importpath = "go-common/library/sync/pipeline",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/net/metadata",
        "//library/time",
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
    srcs = [
        ":package-srcs",
        "//library/sync/pipeline/fanout:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
