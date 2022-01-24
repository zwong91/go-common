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
        "//library/net/rpc/context:all-srcs",
        "//library/net/rpc/interceptor:all-srcs",
        "//library/net/rpc/liverpc:all-srcs",
        "//library/net/rpc/warden:all-srcs",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "rpc",
    srcs = [
        "balancer.go",
        "client.go",
        "client2.go",
        "server.go",
        "trace.go",
    ],
    importpath = "go-common/library/net/rpc",
    tags = ["automanaged"],
    deps = [
        "//library/conf/dsn",
        "//library/conf/env",
        "//library/ecode",
        "//library/log",
        "//library/naming",
        "//library/naming/discovery",
        "//library/net/metadata",
        "//library/net/netutil/breaker",
        "//library/net/rpc/context",
        "//library/net/rpc/interceptor",
        "//library/net/trace",
        "//library/stat",
        "//library/sync/errgroup",
        "//library/time",
        "//vendor/github.com/pkg/errors",
    ],
)

go_test(
    name = "rpc_test",
    srcs = [
        "breaker_test.go",
        "client2_test.go",
        "server_test.go",
    ],
    embed = [":rpc"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/conf/env",
        "//library/ecode",
        "//library/log",
        "//library/naming",
        "//library/naming/discovery",
        "//library/net/rpc/context",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
