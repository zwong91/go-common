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
    data = ["thumbup-test.toml"],
    importpath = "go-common/app/service/main/thumbup/cmd",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/thumbup/conf",
        "//app/service/main/thumbup/server/gorpc",
        "//app/service/main/thumbup/server/grpc",
        "//app/service/main/thumbup/server/http",
        "//app/service/main/thumbup/service",
        "//library/ecode/tip",
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
