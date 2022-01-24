package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "report",
    srcs = [
        "agent.go",
        "conf.go",
    ],
    importpath = "go-common/library/queue/databus/report",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/conf/env",
        "//library/queue/databus",
        "//library/time",
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

go_test(
    name = "report_test",
    srcs = ["agent_test.go"],
    embed = [":report"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/queue/databus",
        "//library/time",
    ],
)
