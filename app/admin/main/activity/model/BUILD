package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "model_test",
    srcs = ["model_test.go"],
    embed = [":model"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "model",
    srcs = [
        "archive.go",
        "bws.go",
        "likes.go",
        "match.go",
        "subject.go",
    ],
    importpath = "go-common/app/admin/main/activity/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/time"],
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
        "//app/admin/main/activity/model/kfc:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
