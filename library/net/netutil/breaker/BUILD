package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "breaker_test",
    srcs = [
        "breaker_test.go",
        "example_test.go",
        "sre_breaker_test.go",
    ],
    embed = [":breaker"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/time",
        "//vendor/github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "breaker",
    srcs = [
        "breaker.go",
        "sre_breaker.go",
    ],
    importpath = "go-common/library/net/netutil/breaker",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/ecode",
        "//library/log",
        "//library/stat/summary",
        "//library/time",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["example_test.go"],
    tags = ["automanaged"],
    deps = [
        "//library/net/netutil/breaker:go_default_library",
        "//library/time:go_default_library",
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
