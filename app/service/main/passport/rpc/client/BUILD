package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "client_test",
    srcs = ["client_test.go"],
    embed = [":client"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//app/service/main/passport/model"],
)

go_library(
    name = "client",
    srcs = ["client.go"],
    importpath = "go-common/app/service/main/passport/rpc/client",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/passport/model",
        "//library/net/rpc",
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
