package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "server_test",
    srcs = ["rpc_test.go"],
    embed = [":server"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/assist/conf",
        "//app/service/main/assist/model/assist",
        "//app/service/main/assist/service",
        "//vendor/github.com/davecgh/go-spew/spew",
    ],
)

go_library(
    name = "server",
    srcs = [
        "rpc.go",
        "rpc_assist.go",
        "rpc_log.go",
    ],
    importpath = "go-common/app/service/main/assist/rpc/server",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/assist/conf",
        "//app/service/main/assist/model/assist",
        "//app/service/main/assist/service",
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
