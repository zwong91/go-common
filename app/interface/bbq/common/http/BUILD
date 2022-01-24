package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "header.go",
        "session.go",
        "uilog.go",
    ],
    importpath = "go-common/app/interface/bbq/common/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/bbq/common/model",
        "//library/log",
        "//library/log/infoc",
        "//library/net/http/blademaster",
        "//library/net/trace",
        "//vendor/github.com/Dai0522/go-hash/murmur3",
        "//vendor/github.com/json-iterator/go",
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
