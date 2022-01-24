package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "grpc",
    srcs = ["server.go"],
    importpath = "go-common/app/service/video/stream-mng/server/grpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/video/stream-mng/api/v1:api",
        "//app/service/video/stream-mng/common",
        "//app/service/video/stream-mng/service",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/net/rpc/warden",
        "//vendor/google.golang.org/grpc",
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
