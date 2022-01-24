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
    srcs = [
        ":package-srcs",
        "//library/cache/memcache/test:all-srcs",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "memcache",
    srcs = [
        "client.go",
        "conn.go",
        "errors.go",
        "memcache.go",
        "mock.go",
        "pool.go",
        "trace.go",
        "util.go",
    ],
    importpath = "go-common/library/cache/memcache",
    tags = ["automanaged"],
    deps = [
        "//library/container/pool",
        "//library/net/trace",
        "//library/stat",
        "//library/time",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/pkg/errors",
    ],
)

go_test(
    name = "memcache_test",
    srcs = [
        "client_test.go",
        "conn_test.go",
        "memcache_test.go",
        "pool_test.go",
        "util_test.go",
    ],
    embed = [":memcache"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/cache/memcache/test",
        "//library/container/pool",
        "//library/time",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/stretchr/testify/assert",
    ],
)
