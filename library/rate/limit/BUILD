package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "limit_test",
    srcs = ["limit_test.go"],
    embed = [":limit"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//library/rate"],
)

go_library(
    name = "limit",
    srcs = ["limit.go"],
    importpath = "go-common/library/rate/limit",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/container/queue/aqm",
        "//library/log",
        "//library/rate",
        "//library/rate/vegas",
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
        "//library/rate/limit/bench/stress/cmd:all-srcs",
        "//library/rate/limit/bench/stress/conf:all-srcs",
        "//library/rate/limit/bench/stress/dao:all-srcs",
        "//library/rate/limit/bench/stress/http:all-srcs",
        "//library/rate/limit/bench/stress/service:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
