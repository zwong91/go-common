package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "actriearea_test",
    srcs = ["actrie_test.go"],
    embed = [":actriearea"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "actriearea",
    srcs = ["actrie.go"],
    importpath = "go-common/app/service/main/filter/model/actriearea",
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
