load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_library(
    name = "cmd_lib",
    srcs = ["main.go"],
    data = ["secure-service-test.toml"],
    importpath = "go-common/app/service/main/secure/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
    deps = [
        "//app/service/main/secure/conf",
        "//app/service/main/secure/http",
        "//app/service/main/secure/rpc/server",
        "//app/service/main/secure/service",
        "//library/log",
        "//library/net/rpc",
        "//library/net/trace",
    ],
)

go_binary(
    name = "cmd",
    embed = [":cmd_lib"],
    visibility = ["//visibility:public"],
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
