package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "configcenter",
    srcs = ["sven.go"],
    importpath = "go-common/app/service/ops/log-agent/conf/configcenter",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/conf",
        "//library/log",
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
