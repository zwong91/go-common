package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "gorpc_test",
    srcs = ["rpc_test.go"],
    embed = [":gorpc"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "gorpc",
    srcs = ["rpc.go"],
    importpath = "go-common/app/service/main/sms/server/gorpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/sms/api",
        "//app/service/main/sms/conf",
        "//app/service/main/sms/model",
        "//app/service/main/sms/service",
        "//library/net/rpc",
        "//library/net/rpc/context",
        "//library/net/rpc/interceptor",
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
