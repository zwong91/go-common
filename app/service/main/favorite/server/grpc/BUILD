package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "grpc",
    srcs = ["server.go"],
    importpath = "go-common/app/service/main/favorite/server/grpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/favorite/api",
        "//app/service/main/favorite/model",
        "//app/service/main/favorite/service",
        "//library/ecode",
        "//library/net/rpc/warden",
        "//vendor/github.com/golang/protobuf/ptypes/empty",
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
