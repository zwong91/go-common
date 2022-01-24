package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "discovery_test",
    srcs = ["discovery_test.go"],
    embed = [":discovery"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/bns/agent/backend",
        "//library/log",
    ],
)

go_library(
    name = "discovery",
    srcs = ["discovery.go"],
    importpath = "go-common/app/service/main/bns/agent/backend/discovery",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/bns/agent/backend",
        "//library/log",
        "//library/naming",
        "//library/stat/prom",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
