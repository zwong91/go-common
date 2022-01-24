package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "bnj",
    srcs = ["bnj.go"],
    importpath = "go-common/app/interface/main/activity/model/bnj",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//app/service/main/archive/api"],
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
