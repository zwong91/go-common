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
    data = ["account-service-example.toml"],
    importpath = "go-common/app/service/main/account/cmd",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/account/conf",
        "//app/service/main/account/rpc/server",
        "//app/service/main/account/server/grpc",
        "//app/service/main/account/server/http",
        "//app/service/main/account/service",
        "//library/log",
        "//library/net/rpc/warden",
        "//library/net/rpc/warden/resolver/livezk",
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
