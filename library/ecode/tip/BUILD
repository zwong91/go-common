package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "tip_test",
    srcs = [
        "example_test.go",
        "tip_test.go",
    ],
    embed = [":tip"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/ecode",
        "//library/net/http/blademaster",
        "//library/net/netutil/breaker",
        "//library/time",
    ],
)

go_library(
    name = "tip",
    srcs = ["tip.go"],
    importpath = "go-common/library/ecode/tip",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/time",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["example_test.go"],
    tags = ["automanaged"],
    deps = [
        "//library/ecode/tip:go_default_library",
        "//library/net/http/blademaster:go_default_library",
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
