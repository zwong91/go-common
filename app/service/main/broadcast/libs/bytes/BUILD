package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "bytes_test",
    srcs = ["buffer_test.go"],
    embed = [":bytes"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "bytes",
    srcs = [
        "buffer.go",
        "writer.go",
    ],
    importpath = "go-common/app/service/main/broadcast/libs/bytes",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
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
