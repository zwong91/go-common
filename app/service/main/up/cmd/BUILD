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
    data = ["up-service.toml"],
    importpath = "go-common/app/service/main/up/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/up/conf",
        "//app/service/main/up/server/gorpc",
        "//app/service/main/up/server/grpc",
        "//app/service/main/up/server/http",
        "//app/service/main/up/service",
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
