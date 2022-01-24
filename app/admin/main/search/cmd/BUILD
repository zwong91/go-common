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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "cmd_lib",
    srcs = ["main.go"],
    data = ["search-admin-test.toml"],
    importpath = "go-common/app/admin/main/search/cmd",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/search/conf",
        "//app/admin/main/search/http",
        "//app/admin/main/search/service",
        "//library/log",
        "//library/net/trace",
    ],
)

go_binary(
    name = "cmd",
    embed = [":cmd_lib"],
    tags = ["automanaged"],
)
