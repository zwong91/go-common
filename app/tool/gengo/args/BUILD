package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "args_test",
    srcs = ["args_test.go"],
    embed = [":args"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//app/tool/gengo/types"],
)

go_library(
    name = "args",
    srcs = [
        "args.go",
        "var.go",
    ],
    importpath = "go-common/app/tool/gengo/args",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/gengo/generator",
        "//app/tool/gengo/namer",
        "//app/tool/gengo/parser",
        "//app/tool/gengo/types",
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
