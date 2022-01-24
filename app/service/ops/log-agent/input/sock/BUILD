package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "sock",
    srcs = [
        "conf.go",
        "sock.go",
    ],
    importpath = "go-common/app/service/ops/log-agent/input/sock",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/ops/log-agent/event",
        "//app/service/ops/log-agent/input",
        "//app/service/ops/log-agent/pkg/flowmonitor",
        "//app/service/ops/log-agent/pkg/lancermonitor",
        "//app/service/ops/log-agent/pkg/lancerroute",
        "//library/log",
        "//library/time",
        "//vendor/github.com/BurntSushi/toml",
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
