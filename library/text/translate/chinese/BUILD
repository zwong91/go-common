package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "chinese_test",
    srcs = ["opencc_test.go"],
    embed = [":chinese"],
    tags = ["automanaged"],
)

go_library(
    name = "chinese",
    srcs = [
        "constant.go",
        "dict.go",
        "opencc.go",
    ],
    importpath = "go-common/library/text/translate/chinese",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//vendor/github.com/go-ego/cedar",
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
