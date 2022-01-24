package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "livezk",
    srcs = ["helper.go"],
    importpath = "go-common/library/net/rpc/warden/resolver/livezk",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/naming",
        "//library/naming/livezk",
        "//library/net/ip",
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
