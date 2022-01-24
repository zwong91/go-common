package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "server_test",
    srcs = [
        "pendant_test.go",
        "rpc_test.go",
    ],
    embed = [":server"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/usersuit/conf",
        "//app/service/main/usersuit/model",
        "//app/service/main/usersuit/service",
        "//library/log",
    ],
)

go_library(
    name = "server",
    srcs = [
        "common.go",
        "medal.go",
        "pendant.go",
        "rpc.go",
    ],
    importpath = "go-common/app/service/main/usersuit/rpc/server",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/usersuit/conf",
        "//app/service/main/usersuit/model",
        "//app/service/main/usersuit/service",
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
