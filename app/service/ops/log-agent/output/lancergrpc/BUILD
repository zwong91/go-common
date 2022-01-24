package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "lancergrpc",
    srcs = [
        "aggr.go",
        "config.go",
        "lancer.go",
        "pool.go",
    ],
    importpath = "go-common/app/service/ops/log-agent/output/lancergrpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/ops/log-agent/event",
        "//app/service/ops/log-agent/output",
        "//app/service/ops/log-agent/output/cache/file",
        "//app/service/ops/log-agent/output/lancergrpc/lancergateway",
        "//app/service/ops/log-agent/pkg/bufio",
        "//app/service/ops/log-agent/pkg/common",
        "//app/service/ops/log-agent/pkg/flowmonitor",
        "//app/service/ops/log-agent/pkg/lancermonitor",
        "//library/log",
        "//library/time",
        "//vendor/github.com/BurntSushi/toml",
        "//vendor/google.golang.org/grpc",
        "//vendor/google.golang.org/grpc/codes",
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
        "//app/service/ops/log-agent/output/lancergrpc/lancergateway:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
