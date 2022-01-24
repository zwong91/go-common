package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "input",
    srcs = ["input.go"],
    importpath = "go-common/app/service/ops/log-agent/input",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/ops/log-agent/event",
        "//library/log",
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
        "//app/service/ops/log-agent/input/file:all-srcs",
        "//app/service/ops/log-agent/input/sock:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
