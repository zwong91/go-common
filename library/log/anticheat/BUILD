package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "anticheat_test",
    srcs = ["anticheat_test.go"],
    embed = [":anticheat"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//library/log/infoc"],
)

go_library(
    name = "anticheat",
    srcs = ["anticheat.go"],
    importpath = "go-common/library/log/anticheat",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//library/log/infoc",
        "//library/net/http/blademaster",
        "//library/net/metadata",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
