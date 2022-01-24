package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "shuffle_test",
    srcs = ["shuffle_test.go"],
    embed = [":shuffle"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "shuffle",
    srcs = ["shuffle.go"],
    importpath = "go-common/app/service/main/bns/lib/shuffle",
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
