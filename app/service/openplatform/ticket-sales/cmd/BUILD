load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_library(
    name = "cmd_lib",
    srcs = ["main.go"],
    data = ["ticket-sales.toml"],
    importpath = "go-common/app/service/openplatform/ticket-sales/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
    deps = [
        "//app/service/openplatform/ticket-sales/conf",
        "//app/service/openplatform/ticket-sales/server/grpc",
        "//app/service/openplatform/ticket-sales/server/http",
        "//app/service/openplatform/ticket-sales/service",
        "//app/service/openplatform/ticket-sales/service/mis",
        "//library/conf/paladin",
        "//library/ecode/tip",
        "//library/log",
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
