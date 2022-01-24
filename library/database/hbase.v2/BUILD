package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "hbase_v2_test",
    srcs = ["hbase_test.go"],
    embed = [":hbase_v2"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/time",
        "//vendor/github.com/tsuna/gohbase/hrpc",
    ],
)

go_library(
    name = "hbase_v2",
    srcs = [
        "config.go",
        "hbase.go",
    ],
    importpath = "go-common/library/database/hbase.v2",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//library/net/trace",
        "//library/stat",
        "//library/time",
        "//vendor/github.com/tsuna/gohbase",
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
