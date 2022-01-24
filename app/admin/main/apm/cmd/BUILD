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
    data = ["apm-admin-test.toml"],
    importpath = "go-common/app/admin/main/apm/cmd",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/apm/conf",
        "//app/admin/main/apm/http",
        "//app/admin/main/apm/service",
        "//library/ecode/tip",
        "//library/log",
        "//library/os/signal",
        "//library/queue/databus/report",
        "//library/syscall",
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
