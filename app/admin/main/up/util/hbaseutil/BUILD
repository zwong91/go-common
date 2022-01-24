package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "hbaseutil_test",
    srcs = ["parser_test.go"],
    embed = [":hbaseutil"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/tsuna/gohbase/hrpc"],
)

go_library(
    name = "hbaseutil",
    srcs = ["parser.go"],
    importpath = "go-common/app/admin/main/up/util/hbaseutil",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//vendor/github.com/tsuna/gohbase/hrpc",
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
