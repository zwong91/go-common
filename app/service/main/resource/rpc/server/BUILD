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
        "//app/service/main/resource/conf",
        "//app/service/main/resource/model",
        "//app/service/main/resource/service",
    ],
)

go_library(
    name = "server",
    srcs = ["server.go"],
    importpath = "go-common/app/service/main/resource/rpc/server",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/resource/conf",
        "//app/service/main/resource/model",
        "//app/service/main/resource/service",
        "//library/ecode",
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
