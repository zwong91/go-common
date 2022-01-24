package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "model",
    srcs = [
        "change_history.go",
        "contract.go",
        "enum.go",
        "main_vip.go",
        "pay_order.go",
        "price_config.go",
        "qr.go",
        "token.go",
        "user_info.go",
        "yst.go",
    ],
    importpath = "go-common/app/service/main/tv/internal/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/account/api",
        "//library/log",
        "//library/time",
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
