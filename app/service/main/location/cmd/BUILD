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
        "convey-test.toml",
        "location-example.toml",
    ],
    importpath = "go-common/app/service/main/location/cmd",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/location/conf",
        "//app/service/main/location/rpc/server",
        "//app/service/main/location/server/grpc",
        "//app/service/main/location/server/http",
        "//app/service/main/location/service",
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
