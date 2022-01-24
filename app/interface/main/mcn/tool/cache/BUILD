package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "cache",
    srcs = ["cache.go"],
    importpath = "go-common/app/interface/main/mcn/tool/cache",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/cache",
        "//library/cache/memcache",
        "//library/log",
        "//library/net/metadata",
        "//library/stat/prom",
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
