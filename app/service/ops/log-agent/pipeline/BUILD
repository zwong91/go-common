package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "pipeline",
    srcs = [
        "management.go",
        "pipeline.go",
    ],
    importpath = "go-common/app/service/ops/log-agent/pipeline",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/ops/log-agent/event",
        "//app/service/ops/log-agent/input",
        "//app/service/ops/log-agent/input/file",
        "//app/service/ops/log-agent/input/sock",
        "//app/service/ops/log-agent/output",
        "//app/service/ops/log-agent/output/lancergrpc",
        "//app/service/ops/log-agent/output/lancerlogstream",
        "//app/service/ops/log-agent/output/stdout",
        "//app/service/ops/log-agent/pkg/common",
        "//app/service/ops/log-agent/processor",
        "//app/service/ops/log-agent/processor/classify",
        "//app/service/ops/log-agent/processor/fileLog",
        "//app/service/ops/log-agent/processor/grok",
        "//app/service/ops/log-agent/processor/httpstream",
        "//app/service/ops/log-agent/processor/jsonLog",
        "//app/service/ops/log-agent/processor/lengthCheck",
        "//app/service/ops/log-agent/processor/sample",
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
    srcs = [
        ":package-srcs",
        "//app/service/ops/log-agent/pipeline/dockerlogcollector:all-srcs",
        "//app/service/ops/log-agent/pipeline/hostlogcollector:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
