package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "feature",
    srcs = ["feature_gate.go"],
    importpath = "go-common/library/exp/feature",
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

go_test(
    name = "feature_test",
    srcs = [
        "example_test.go",
        "feature_gate_test.go",
    ],
    embed = [":feature"],
    rundir = ".",
    tags = ["automanaged"],
)

go_test(
    name = "go_default_test",
    srcs = ["example_test.go"],
    tags = ["automanaged"],
    deps = ["//library/exp/feature:go_default_library"],
)
