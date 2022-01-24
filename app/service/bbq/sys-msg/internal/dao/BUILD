package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "dao.cache.go",
        "dao.go",
    ],
    importpath = "go-common/app/service/bbq/sys-msg/internal/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/sys-msg/api/v1:api",
        "//app/service/bbq/sys-msg/internal/conf",
        "//library/cache",
        "//library/cache/redis",
        "//library/database/sql",
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
