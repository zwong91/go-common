package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "resolver",
    srcs = [
        "resolver.go",
        "util.go",
    ],
    importpath = "go-common/library/net/rpc/warden/resolver",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/conf/env",
        "//library/log",
        "//library/naming",
        "//library/net/rpc/warden/metadata",
        "//vendor/github.com/dgryski/go-farm",
        "//vendor/github.com/pkg/errors",
        "//vendor/google.golang.org/grpc/resolver",
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
    srcs = [
        ":package-srcs",
        "//library/net/rpc/warden/resolver/direct:all-srcs",
        "//library/net/rpc/warden/resolver/livezk:all-srcs",
        "//library/net/rpc/warden/resolver/test:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
