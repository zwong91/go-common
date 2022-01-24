package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "cache.go",
        "rating.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/up-rating/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/up-rating/conf",
        "//app/interface/main/up-rating/dao",
        "//app/interface/main/up-rating/model",
        "//library/ecode",
        "//library/log",
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
