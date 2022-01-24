package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
    "go_test",
)

go_test(
    name = "bgr_test",
    srcs = ["bgl_test.go"],
    embed = [":bgr_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "bgr_lib",
    srcs = [
        "ast_inspect.go",
        "cache.go",
        "file_filter.go",
        "lint_register.go",
        "main.go",
        "model.go",
        "script_parser.go",
    ],
    importpath = "go-common/app/tool/bgr",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/bgr/log",
        "//vendor/github.com/pkg/errors",
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
        "//app/tool/bgr/log:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_binary(
    name = "bgr",
    embed = [":bgr_lib"],
    tags = ["automanaged"],
)
