package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "proxy_test",
    srcs = [
        "example_test.go",
        "proxy_test.go",
    ],
    embed = [":proxy"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "proxy",
    srcs = ["proxy.go"],
    importpath = "go-common/library/net/http/blademaster/middleware/proxy",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/conf/env",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//vendor/github.com/pkg/errors",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["example_test.go"],
    tags = ["automanaged"],
    deps = [
        "//library/net/http/blademaster:go_default_library",
        "//library/net/http/blademaster/middleware/proxy:go_default_library",
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
