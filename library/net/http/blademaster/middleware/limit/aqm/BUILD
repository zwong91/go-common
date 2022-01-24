package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "aqm_test",
    srcs = [
        "aqm_test.go",
        "example_test.go",
    ],
    embed = [":aqm"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/log",
        "//library/net/http/blademaster",
    ],
)

go_library(
    name = "aqm",
    srcs = ["aqm.go"],
    importpath = "go-common/library/net/http/blademaster/middleware/limit/aqm",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/container/queue/aqm",
        "//library/ecode",
        "//library/net/http/blademaster",
        "//library/rate",
        "//library/rate/limit",
        "//library/stat/prom",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["example_test.go"],
    tags = ["automanaged"],
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
