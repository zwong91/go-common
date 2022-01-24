package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "cmd",
    embed = [":cmd_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "cmd_lib",
    srcs = ["main.go"],
    importpath = "go-common/app/service/ops/log-agent/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/ops/log-agent/conf",
        "//app/service/ops/log-agent/pipeline",
        "//app/service/ops/log-agent/pipeline/dockerlogcollector",
        "//app/service/ops/log-agent/pipeline/hostlogcollector",
        "//app/service/ops/log-agent/pkg/flowmonitor",
        "//app/service/ops/log-agent/pkg/httpstream",
        "//app/service/ops/log-agent/pkg/lancermonitor",
        "//app/service/ops/log-agent/pkg/lancerroute",
        "//app/service/ops/log-agent/pkg/limit",
        "//library/conf/env",
        "//library/log",
        "//library/naming",
        "//library/naming/discovery",
        "//library/net/ip",
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
