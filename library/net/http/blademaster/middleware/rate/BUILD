package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "rate_test",
    srcs = [
        "example_test.go",
        "limit_test.go",
    ],
    embed = [":rate"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//library/net/http/blademaster"],
)

go_library(
    name = "rate",
    srcs = ["limit.go"],
    importpath = "go-common/library/net/http/blademaster/middleware/rate",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/golang.org/x/time/rate",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["example_test.go"],
    tags = ["automanaged"],
    deps = [
        "//library/net/http/blademaster:go_default_library",
        "//library/net/http/blademaster/middleware/rate:go_default_library",
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
