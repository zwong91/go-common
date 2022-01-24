package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "goparser_test",
    srcs = ["goparser_test.go"],
    embed = [":goparser"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "goparser",
    srcs = ["goparser.go"],
    importpath = "go-common/app/tool/warden/goparser",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//app/tool/warden/types"],
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
