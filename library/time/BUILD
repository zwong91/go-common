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
    name = "time",
    srcs = ["time.go"],
    importpath = "go-common/library/time",
    tags = ["automanaged"],
)

go_test(
    name = "time_test",
    srcs = ["time_test.go"],
    embed = [":time"],
    rundir = ".",
    tags = ["automanaged"],
)
