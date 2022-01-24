package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "grpc_test",
    srcs = ["server_test.go"],
    embed = [":grpc"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/location/api",
        "//app/service/main/location/conf",
        "//app/service/main/location/service",
        "//library/net/rpc/warden",
        "//library/time",
    ],
)

go_library(
    name = "grpc",
    srcs = ["server.go"],
    importpath = "go-common/app/service/main/location/server/grpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/location/api",
        "//app/service/main/location/model",
        "//app/service/main/location/service",
        "//library/net/rpc/warden",
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
