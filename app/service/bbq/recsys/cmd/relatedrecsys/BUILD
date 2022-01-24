package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "relatedrecsys",
    embed = [":relatedrecsys_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "relatedrecsys_lib",
    srcs = ["main.go"],
    importpath = "go-common/app/service/bbq/recsys/cmd/relatedrecsys",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/recsys/api/grpc/v1:grpc",
        "//library/net/rpc/warden",
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
