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
    data = ["tag.toml"],
    importpath = "go-common/app/service/main/tag/cmd",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/tag/conf",
        "//app/service/main/tag/rpc/server",
        "//app/service/main/tag/server/grpc",
        "//app/service/main/tag/server/http",
        "//app/service/main/tag/service",
        "//library/ecode/tip",
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
