package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "log_test",
    srcs = [
        "agent_test.go",
        "dsn_test.go",
        "encode_test.go",
        "example_test.go",
        "log_test.go",
        "pattern_test.go",
    ],
    embed = [":log"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/log/internal",
        "//library/net/metadata",
        "//library/time",
        "//vendor/github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "log",
    srcs = [
        "agent.go",
        "doc.go",
        "dsn.go",
        "file.go",
        "handler.go",
        "level.go",
        "log.go",
        "logrus.go",
        "pattern.go",
        "stdout.go",
        "util.go",
        "verbose.go",
    ],
    importpath = "go-common/library/log",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/conf/dsn",
        "//library/conf/env",
        "//library/log/internal",
        "//library/log/internal/filewriter",
        "//library/net/metadata",
        "//library/net/trace",
        "//library/stat/prom",
        "//library/time",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/sirupsen/logrus",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["example_test.go"],
    tags = ["automanaged"],
    deps = ["//library/log:go_default_library"],
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
        "//library/log/anticheat:all-srcs",
        "//library/log/benchmark:all-srcs",
        "//library/log/infoc:all-srcs",
        "//library/log/internal:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
