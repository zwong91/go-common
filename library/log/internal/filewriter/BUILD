package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "filewriter_test",
    srcs = ["filewriter_test.go"],
    embed = [":filewriter"],
    tags = ["automanaged"],
    deps = ["//vendor/github.com/stretchr/testify/assert"],
)

go_library(
    name = "filewriter",
    srcs = [
        "filewriter.go",
        "option.go",
    ],
    importpath = "go-common/library/log/internal/filewriter",
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
