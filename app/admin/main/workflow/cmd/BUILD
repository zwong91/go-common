package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "cmd",
    embed = [":cmd_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "cmd_lib",
    srcs = ["main.go"],
    data = ["workflow-admin-develop.toml"],
    importpath = "go-common/app/admin/main/workflow/cmd",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/workflow/http",
        "//app/admin/main/workflow/service",
        "//library/conf/paladin",
        "//library/ecode/tip",
        "//library/log",
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
