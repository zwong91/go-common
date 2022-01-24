load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "server",
    srcs = [
        "identify.go",
        "rpc.go",
    ],
    importpath = "go-common/app/service/main/identify-game/rpc/server",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/identify-game/conf",
        "//app/service/main/identify-game/model",
        "//app/service/main/identify-game/service",
        "//library/net/rpc",
        "//library/net/rpc/context",
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
