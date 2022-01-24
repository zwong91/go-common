package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "server_test",
    srcs = ["server_test.go"],
    embed = [":server"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/location/conf",
        "//app/service/main/location/model",
        "//app/service/main/location/service",
    ],
)

go_library(
    name = "server",
    srcs = ["server.go"],
    importpath = "go-common/app/service/main/location/rpc/server",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/location/conf",
        "//app/service/main/location/model",
        "//app/service/main/location/service",
        "//library/net/rpc",
        "//library/net/rpc/context",
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
