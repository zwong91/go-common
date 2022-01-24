package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = ["dao.go"],
    importpath = "go-common/app/service/bbq/push/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/push/conf",
        "//app/service/bbq/push/dao/jpush",
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
        "//app/service/bbq/push/dao/jpush:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
