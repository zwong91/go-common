package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "grpc",
    srcs = ["server.go"],
    importpath = "go-common/app/service/live/resource/server/grpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/resource/api/grpc/v1:grpc",
        "//app/service/live/resource/api/grpc/v2:grpc",
        "//app/service/live/resource/conf",
        "//app/service/live/resource/service/v1:service",
        "//app/service/live/resource/service/v2:service",
        "//library/net/rpc/warden",
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
