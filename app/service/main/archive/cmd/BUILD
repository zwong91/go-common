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
    data = [
        "archive-service-example.toml",
        "archive-service-test.toml",
    ],
    importpath = "go-common/app/service/main/archive/cmd",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/conf",
        "//app/service/main/archive/server/gorpc",
        "//app/service/main/archive/server/grpc",
        "//app/service/main/archive/server/http",
        "//app/service/main/archive/service",
        "//library/ecode/tip",
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
