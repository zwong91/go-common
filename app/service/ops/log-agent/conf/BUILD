package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "conf",
    srcs = ["conf.go"],
    importpath = "go-common/app/service/ops/log-agent/conf",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/ops/log-agent/conf/configcenter",
        "//app/service/ops/log-agent/pipeline/dockerlogcollector",
        "//app/service/ops/log-agent/pipeline/hostlogcollector",
        "//app/service/ops/log-agent/pkg/flowmonitor",
        "//app/service/ops/log-agent/pkg/httpstream",
        "//app/service/ops/log-agent/pkg/lancermonitor",
        "//app/service/ops/log-agent/pkg/limit",
        "//library/log",
        "//library/naming/discovery",
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
        "//app/service/ops/log-agent/conf/configcenter:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
