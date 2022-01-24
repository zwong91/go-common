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
    data = ["broadcast.toml"],
    importpath = "go-common/app/interface/main/broadcast/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/broadcast/conf",
        "//app/interface/main/broadcast/server",
        "//app/interface/main/broadcast/server/grpc",
        "//app/interface/main/broadcast/server/http",
        "//library/conf/env",
        "//library/conf/paladin",
        "//library/log",
        "//library/naming",
        "//library/naming/discovery",
        "//library/net/ip",
        "//library/net/rpc/warden/resolver",
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
