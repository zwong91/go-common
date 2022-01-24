package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "testdata",
    srcs = ["client.go"],
    importpath = "go-common/library/net/rpc/liverpc/testdata",
    tags = ["manual"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/net/rpc/liverpc",
        "//library/net/rpc/liverpc/testdata/v1:testdata",
        "//library/net/rpc/liverpc/testdata/v2:testdata",
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
    srcs = [
        ":package-srcs",
        "//library/net/rpc/liverpc/testdata/v1:all-srcs",
        "//library/net/rpc/liverpc/testdata/v2:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
