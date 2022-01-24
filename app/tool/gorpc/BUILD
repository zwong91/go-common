load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

package(default_visibility = ["//visibility:public"])

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
        "//app/tool/gorpc/goparser:all-srcs",
        "//app/tool/gorpc/input:all-srcs",
        "//app/tool/gorpc/model:all-srcs",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "gorpc_lib",
    srcs = ["main.go"],
    importpath = "go-common/app/tool/gorpc",
    tags = ["automanaged"],
    deps = [
        "//app/tool/gorpc/goparser",
        "//app/tool/gorpc/input",
        "//app/tool/gorpc/model",
        "//vendor/golang.org/x/tools/imports",
    ],
)

go_binary(
    name = "gorpc",
    embed = [":gorpc_lib"],
)
