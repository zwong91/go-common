package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "backend_test",
    srcs = ["backend_test.go"],
    embed = [":backend"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "backend",
    srcs = ["backend.go"],
    importpath = "go-common/app/service/main/bns/agent/backend",
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
    srcs = [
        ":package-srcs",
        "//app/service/main/bns/agent/backend/discovery:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
