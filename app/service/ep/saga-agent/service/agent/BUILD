package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "agent_test",
    srcs = ["agent_test.go"],
    embed = [":agent"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "agent",
    srcs = ["agent.go"],
    data = ["runners.toml"],
    importpath = "go-common/app/service/ep/saga-agent/service/agent",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/ep/saga-agent/conf",
        "//library/log",
        "//vendor/github.com/pkg/errors",
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
