package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "authority.go",
        "cache.go",
        "parameter.go",
        "score.go",
        "service.go",
        "statistics.go",
        "transfer.go",
    ],
    importpath = "go-common/app/admin/main/up-rating/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up-rating/conf",
        "//app/admin/main/up-rating/dao",
        "//app/admin/main/up-rating/dao/global",
        "//app/admin/main/up-rating/model",
        "//library/ecode",
        "//library/log",
        "//library/xstr",
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
