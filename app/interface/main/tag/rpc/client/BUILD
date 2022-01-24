package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "client_test",
    srcs = [
        "tag2_test.go",
        "tag_test.go",
    ],
    embed = [":client"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tag/model",
        "//library/net/rpc",
    ],
)

go_library(
    name = "client",
    srcs = [
        "channel.go",
        "tag.go",
        "tag2.go",
    ],
    importpath = "go-common/app/interface/main/tag/rpc/client",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tag/model",
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
