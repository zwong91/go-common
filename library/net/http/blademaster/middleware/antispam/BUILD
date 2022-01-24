package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "antispam_test",
    srcs = [
        "antispam_test.go",
        "example_test.go",
    ],
    embed = [":antispam"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/cache/redis",
        "//library/container/pool",
        "//library/net/http/blademaster",
        "//library/time",
        "//vendor/github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "antispam",
    srcs = ["antispam.go"],
    importpath = "go-common/library/net/http/blademaster/middleware/antispam",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/cache/redis",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/github.com/pkg/errors",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["example_test.go"],
    tags = ["automanaged"],
    deps = [
        "//library/cache/redis:go_default_library",
        "//library/container/pool:go_default_library",
        "//library/net/http/blademaster:go_default_library",
        "//library/net/http/blademaster/middleware/antispam:go_default_library",
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
