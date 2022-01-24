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
    data = ["sms-service-test.toml"],
    importpath = "go-common/app/service/main/sms/cmd",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/sms/conf",
        "//app/service/main/sms/server/gorpc",
        "//app/service/main/sms/server/grpc",
        "//app/service/main/sms/server/http",
        "//app/service/main/sms/service",
        "//library/log",
        "//library/net/trace",
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
