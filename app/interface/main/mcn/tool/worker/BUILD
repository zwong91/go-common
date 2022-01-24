package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "worker",
    srcs = ["pool.go"],
    importpath = "go-common/app/interface/main/mcn/tool/worker",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/log"],
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

go_test(
    name = "worker_test",
    srcs = ["pool_test.go"],
    embed = [":worker"],
    rundir = ".",
    tags = ["automanaged"],
)
