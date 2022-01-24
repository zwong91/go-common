package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "liverpcgen",
    embed = [":liverpcgen_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "liverpcgen_lib",
    srcs = ["main.go"],
    importpath = "go-common/app/tool/liverpc/liverpcgen",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/siddontang/go/ioutil2",
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
