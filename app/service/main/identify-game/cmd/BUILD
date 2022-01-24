load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_library(
    name = "cmd_lib",
    srcs = ["main.go"],
    data = ["identify-game-service.toml"],
    importpath = "go-common/app/service/main/identify-game/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
    deps = [
        "//app/service/main/identify-game/conf",
        "//app/service/main/identify-game/rpc/server",
        "//app/service/main/identify-game/server/grpc",
        "//app/service/main/identify-game/server/http",
        "//app/service/main/identify-game/service",
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
