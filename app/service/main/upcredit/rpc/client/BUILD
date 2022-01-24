package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "client_test",
    srcs = ["up_test.go"],
    embed = [":client"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "client",
    srcs = ["up.go"],
    importpath = "go-common/app/service/main/upcredit/rpc/client",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/net/rpc"],
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
