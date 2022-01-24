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
        "article-test.toml",
        "convey-test.toml",
    ],
    importpath = "go-common/app/interface/openplatform/article/cmd",
    tags = ["automanaged"],
    deps = [
        "//app/interface/openplatform/article/conf",
        "//app/interface/openplatform/article/http",
        "//app/interface/openplatform/article/rpc/server",
        "//app/interface/openplatform/article/service",
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
