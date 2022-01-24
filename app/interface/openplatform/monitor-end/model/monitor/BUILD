package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "monitor",
    srcs = [
        "handler.go",
        "monitor.go",
    ],
    importpath = "go-common/app/interface/openplatform/monitor-end/model/monitor",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/conf/env",
        "//library/log",
        "//library/net/trace",
        "//library/time",
        "//vendor/github.com/json-iterator/go",
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
    name = "monitor_test",
    srcs = [
        "handler_test.go",
        "monitor_test.go",
    ],
    embed = [":monitor"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
