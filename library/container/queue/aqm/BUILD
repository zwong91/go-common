package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "aqm_test",
    srcs = ["codel_test.go"],
    embed = [":aqm"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//library/ecode"],
)

go_library(
    name = "aqm",
    srcs = ["codel.go"],
    importpath = "go-common/library/container/queue/aqm",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/ecode"],
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
