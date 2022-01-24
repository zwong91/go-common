package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "counter",
    srcs = [
        "collector.go",
        "counter.go",
        "desc.go",
        "fnv.go",
        "go_collector.go",
        "metric.go",
        "observer.go",
        "process_collector.go",
        "registry.go",
        "summary.go",
        "untyped.go",
        "value.go",
        "vec.go",
    ],
    importpath = "go-common/app/service/ops/log-agent/pkg/flowmonitor/counter",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/beorn7/perks/quantile",
        "//vendor/github.com/golang/protobuf/proto",
        "//vendor/github.com/prometheus/client_model/go",
        "//vendor/github.com/prometheus/common/model",
        "//vendor/github.com/prometheus/procfs",
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
