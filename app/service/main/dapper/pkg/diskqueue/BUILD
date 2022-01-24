package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "diskqueue_test",
    srcs = [
        "bucket_test.go",
        "diskqueue_test.go",
    ],
    embed = [":diskqueue"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "diskqueue",
    srcs = [
        "bucket.go",
        "diskqueue.go",
    ],
    importpath = "go-common/app/service/main/dapper/pkg/diskqueue",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
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
