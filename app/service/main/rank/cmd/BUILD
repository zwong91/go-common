package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
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

go_binary(
    name = "cmd",
    embed = [":cmd_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "cmd_lib",
    srcs = ["main.go"],
    data = [
        "rank-service.toml",
        "test.toml",
    ],
    importpath = "go-common/app/service/main/rank/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/rank/conf",
        "//app/service/main/rank/server/gorpc",
        "//app/service/main/rank/server/http",
        "//app/service/main/rank/service",
        "//library/conf/env",
        "//library/ecode/tip",
        "//library/log",
        "//library/naming",
        "//library/naming/discovery",
        "//library/net/ip",
        "//library/net/trace",
    ],
)
