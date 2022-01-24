package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "server_test",
    srcs = ["rpcserver_test.go"],
    embed = [":server"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/upcredit/conf",
        "//app/service/main/upcredit/service",
        "//library/net/rpc",
        "//library/time",
    ],
)

go_library(
    name = "server",
    srcs = ["rpcserver.go"],
    importpath = "go-common/app/service/main/upcredit/rpc/server",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/upcredit/conf",
        "//app/service/main/upcredit/service",
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
