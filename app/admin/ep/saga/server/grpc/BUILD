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
        "//app/admin/ep/saga/api/grpc/v1:grpc",
        "//library/net/rpc/warden",
        "//library/net/rpc/warden/resolver",
        "//library/net/rpc/warden/resolver/direct",
    ],
)

go_library(
    name = "grpc",
    srcs = ["server.go"],
    importpath = "go-common/app/admin/ep/saga/server/grpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/ep/saga/api/grpc/v1:grpc",
        "//app/admin/ep/saga/service/wechat",
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
