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
    data = ["tag-example.toml"],
    importpath = "go-common/app/interface/main/tag/cmd",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tag/conf",
        "//app/interface/main/tag/http",
        "//app/interface/main/tag/rpc/server",
        "//app/interface/main/tag/server/grpc",
        "//app/interface/main/tag/service",
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
