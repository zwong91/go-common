package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "client_test",
    srcs = ["feed_test.go"],
    embed = [":client"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//app/service/main/feed/model"],
)

go_library(
    name = "client",
    srcs = ["feed.go"],
    importpath = "go-common/app/service/main/feed/rpc/client",
    tags = ["automanaged"],
    deps = [
        "//app/interface/openplatform/article/model",
        "//app/service/main/feed/model",
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
