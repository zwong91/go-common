load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

package(default_visibility = ["//visibility:public"])

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
)

go_library(
    name = "redis",
    srcs = [
        "commandinfo.go",
        "conn.go",
        "doc.go",
        "errors.go",
        "log.go",
        "mock.go",
        "pool.go",
        "pubsub.go",
        "redis.go",
        "reply.go",
        "scan.go",
        "script.go",
        "trace.go",
    ],
    importpath = "go-common/library/cache/redis",
    tags = ["automanaged"],
    deps = [
        "//library/container/pool",
        "//library/net/trace",
        "//library/stat",
        "//library/time",
        "//vendor/github.com/pkg/errors",
    ],
)

go_test(
    name = "redis_test",
    srcs = [
        "commandinfo_test.go",
        "conn_test.go",
        "pool_test.go",
        "pubsub_test.go",
        "redis_test.go",
        "reply_test.go",
        "scan_test.go",
        "script_test.go",
        "test_test.go",
        "trace_test.go",
    ],
    embed = [":redis"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/container/pool",
        "//library/net/trace",
        "//library/time",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/stretchr/testify/assert",
    ],
)
