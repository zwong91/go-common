load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "client",
    srcs = ["secure.go"],
    importpath = "go-common/app/service/main/secure/rpc/client",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/secure/model",
        "//library/net/rpc",
    ],
)

go_test(
    name = "client_test",
    srcs = ["secure_test.go"],
    embed = [":client"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//app/service/main/secure/model"],
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
