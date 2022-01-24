package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "account",
    srcs = ["account.go"],
    importpath = "go-common/app/service/openplatform/ticket-sales/model/order_checker/account",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/account/api",
        "//app/service/main/account/model",
        "//app/service/main/vip/model",
        "//app/service/openplatform/ticket-sales/api/grpc/v1:grpc",
        "//app/service/openplatform/ticket-sales/dao",
        "//library/ecode",
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
