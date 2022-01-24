package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "livezk_test",
    srcs = ["livezk_test.go"],
    embed = [":livezk"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/naming",
        "//library/time",
    ],
)

go_library(
    name = "livezk",
    srcs = ["livezk.go"],
    importpath = "go-common/library/naming/livezk",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//library/naming",
        "//library/time",
        "//vendor/github.com/samuel/go-zookeeper/zk",
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
