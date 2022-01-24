package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "cmd",
    embed = [":cmd_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "cmd_lib",
    srcs = ["main.go"],
    data = [
        "dynamic-service-example.toml",
        "dynamic-service-test.toml",
    ],
    importpath = "go-common/app/service/main/dynamic/cmd",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/dynamic/conf",
        "//app/service/main/dynamic/http",
        "//app/service/main/dynamic/rpc/server",
        "//app/service/main/dynamic/service",
        "//library/log",
        "//library/net/trace",
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
