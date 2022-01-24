package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "client",
    srcs = [
        "mock.go",
        "thumbup.go",
    ],
    importpath = "go-common/app/service/main/thumbup/rpc/client",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/thumbup/model",
        "//library/net/rpc",
        "//vendor/github.com/golang/mock/gomock",
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
