package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "script",
    embed = [":script_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "script_lib",
    srcs = ["main.go"],
    importpath = "go-common/app/admin/main/cache/script",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/gopkg.in/yaml.v2:yaml_v2"],
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
