package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = ["http.go"],
    importpath = "go-common/library/rate/limit/bench/stress/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//library/net/http/blademaster",
        "//library/rate",
        "//library/rate/limit",
        "//library/rate/limit/bench/stress/conf",
        "//library/rate/limit/bench/stress/service",
        "//library/rate/vegas",
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
