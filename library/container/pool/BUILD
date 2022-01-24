package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "pool_test",
    srcs = [
        "list_test.go",
        "slice_test.go",
    ],
    embed = [":pool"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/time",
        "//vendor/github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "pool",
    srcs = [
        "list.go",
        "pool.go",
        "slice.go",
    ],
    importpath = "go-common/library/container/pool",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/time"],
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
