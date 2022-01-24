package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "model",
    srcs = [
        "bugly.go",
        "constants.go",
        "marthe.go",
        "model.go",
        "tapd.go",
        "wechat.go",
    ],
    importpath = "go-common/app/admin/ep/marthe/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/ecode"],
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
