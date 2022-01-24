package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "classify",
    srcs = [
        "classify.go",
        "config.go",
    ],
    importpath = "go-common/app/service/ops/log-agent/processor/classify",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/ops/log-agent/conf/configcenter",
        "//app/service/ops/log-agent/event",
        "//app/service/ops/log-agent/pkg/common",
        "//app/service/ops/log-agent/processor",
        "//library/log",
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
