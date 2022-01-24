package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "lrulist_test",
    srcs = ["list_test.go"],
    embed = [":lrulist"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "lrulist",
    srcs = ["list.go"],
    importpath = "go-common/app/service/main/filter/model/lrulist",
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
