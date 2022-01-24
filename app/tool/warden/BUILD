package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "warden",
    embed = [":warden_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "warden_lib",
    srcs = ["main.go"],
    importpath = "go-common/app/tool/warden",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/warden/generator",
        "//app/tool/warden/goparser",
        "//app/tool/warden/types",
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
    srcs = [
        ":package-srcs",
        "//app/tool/warden/generator:all-srcs",
        "//app/tool/warden/goparser:all-srcs",
        "//app/tool/warden/types:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
