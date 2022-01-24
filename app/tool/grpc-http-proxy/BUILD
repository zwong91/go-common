package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_library(
    name = "grpc-http-proxy_lib",
    srcs = [
        "grpc.go",
        "http.go",
        "main.go",
    ],
    importpath = "go-common/app/tool/grpc-http-proxy",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/ecode",
        "//library/log",
        "//library/naming/discovery",
        "//library/net/http/blademaster",
        "//library/time",
        "//vendor/github.com/gogo/protobuf/jsonpb",
        "//vendor/google.golang.org/grpc",
        "//vendor/google.golang.org/grpc/encoding",
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

go_binary(
    name = "grpc-http-proxy",
    embed = [":grpc-http-proxy_lib"],
    tags = ["automanaged"],
)
