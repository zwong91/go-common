package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "api.go",
        "dao.go",
        "token.go",
    ],
    importpath = "go-common/app/admin/live/live-admin/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/live/live-admin/conf",
        "//app/service/live/av/api/liverpc",
        "//app/service/live/relation/api/liverpc",
        "//library/cache/redis",
        "//library/log",
        "//library/net/rpc/liverpc",
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
