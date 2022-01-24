package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "benchmark",
    embed = [":benchmark_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "benchmark_lib",
    srcs = ["benchmark.go"],
    importpath = "go-common/app/service/main/riot-search/benchmark",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/riot-search/model",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/time",
        "//vendor/github.com/ivpusic/grpool",
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
