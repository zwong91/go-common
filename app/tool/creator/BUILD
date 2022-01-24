package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_library(
    name = "creator_lib",
    srcs = [
        "gen.go",
        "main.go",
        "parser.go",
        "templete.go",
        "upgrade.go",
        "utils.go",
    ],
    importpath = "go-common/app/tool/creator",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/otokaze/mock/mockgen",
        "//vendor/github.com/otokaze/mock/mockgen/model",
        "//vendor/golang.org/x/tools/imports",
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

go_binary(
    name = "creator",
    embed = [":creator_lib"],
    tags = ["automanaged"],
)
