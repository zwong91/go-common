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
    data = ["history-interface-test.toml"],
    importpath = "go-common/app/interface/main/history/cmd",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/history/conf",
        "//app/interface/main/history/http",
        "//app/interface/main/history/server/gorpc",
        "//app/interface/main/history/server/grpc",
        "//app/interface/main/history/service",
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
