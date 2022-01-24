load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "cmd",
    embed = [":cmd_lib"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "cmd_lib",
    srcs = ["main.go"],
    data = ["feed-admin-test.toml"],
    importpath = "go-common/app/admin/main/feed/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/feed/conf",
        "//app/admin/main/feed/http",
        "//library/log",
        "//library/net/trace",
        "//library/queue/databus/report",
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
