load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

package(default_visibility = ["//visibility:public"])

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
)

go_library(
    name = "metadata",
    srcs = [
        "key.go",
        "metadata.go",
    ],
    importpath = "go-common/library/net/metadata",
    tags = ["automanaged"],
)

go_test(
    name = "metadata_test",
    srcs = ["metadata_test.go"],
    embed = [":metadata"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/stretchr/testify/assert"],
)
