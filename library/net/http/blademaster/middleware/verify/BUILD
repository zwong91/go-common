package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "verify_test",
    srcs = [
        "example_test.go",
        "verify_test.go",
    ],
    embed = [":verify"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/net/netutil/breaker",
        "//library/time",
        "//vendor/github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "verify",
    srcs = ["verify.go"],
    importpath = "go-common/library/net/http/blademaster/middleware/verify",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/time",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["example_test.go"],
    tags = ["automanaged"],
    deps = [
        "//library/net/http/blademaster:go_default_library",
        "//library/net/http/blademaster/middleware/verify:go_default_library",
        "//library/net/metadata:go_default_library",
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
