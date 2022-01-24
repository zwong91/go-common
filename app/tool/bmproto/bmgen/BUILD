package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "bmgen",
    embed = [":bmgen_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "bmgen_lib",
    srcs = ["main.go"],
    importpath = "go-common/app/tool/bmproto/bmgen",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/urfave/cli",
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
