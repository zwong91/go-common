package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "model",
    srcs = [
        "archive.go",
        "const.go",
        "model.go",
        "retry.go",
    ],
    importpath = "go-common/app/job/main/app/model",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/app/model/space",
        "//library/conf/env",
        "//library/time",
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
        "//app/job/main/app/model/feed:all-srcs",
        "//app/job/main/app/model/space:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
