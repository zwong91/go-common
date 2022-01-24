package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = ["api.go"],
    importpath = "go-common/app/interface/live/live-demo/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/live-demo/conf",
        "//app/service/live/room/api/liverpc",
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
    srcs = [
        ":package-srcs",
        "//app/interface/live/live-demo/dao/foo:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
