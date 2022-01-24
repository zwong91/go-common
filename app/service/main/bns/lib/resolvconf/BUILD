package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "resolvconf_test",
    srcs = ["resolvconf_unix_test.go"],
    embed = [":resolvconf"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "resolvconf",
    srcs = [
        "resolvconf_unix.go",
        "resolvconf_windows.go",
    ],
    importpath = "go-common/app/service/main/bns/lib/resolvconf",
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
