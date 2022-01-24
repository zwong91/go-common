package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "agent",
    embed = [":agent_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "agent_lib",
    srcs = ["main.go"],
    data = [
        "dapper-agent.toml",
        "dapper-agent-example.toml",
    ],
    importpath = "go-common/app/service/main/dapper/cmd/agent",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/dapper/agent",
        "//app/service/main/dapper/conf",
        "//app/service/main/dapper/pkg/util",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
