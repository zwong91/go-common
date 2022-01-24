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
    data = ["config-service-example.toml"],
    importpath = "go-common/app/infra/config/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/infra/config/conf",
        "//app/infra/config/http",
        "//app/infra/config/rpc/server",
        "//app/infra/config/service/v1:service",
        "//app/infra/config/service/v2:service",
        "//library/conf/env",
        "//library/log",
        "//library/naming",
        "//library/naming/discovery",
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
