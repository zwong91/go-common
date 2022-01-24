package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "time_test",
    srcs = ["timer_test.go"],
    embed = [":time"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//library/log"],
)

go_library(
    name = "time",
    srcs = [
        "debug.go",
        "timer.go",
    ],
    importpath = "go-common/app/service/main/broadcast/libs/time",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/log"],
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
