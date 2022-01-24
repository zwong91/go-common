package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "blademaster_test",
    srcs = [
        "client_test.go",
        "example_test.go",
        "server_test.go",
        "trace_test.go",
    ],
    embed = [":blademaster"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/http/blademaster/render",
        "//library/net/http/blademaster/tests",
        "//library/net/metadata",
        "//library/net/netutil/breaker",
        "//library/net/trace",
        "//library/time",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "blademaster",
    srcs = [
        "client.go",
        "context.go",
        "cors.go",
        "csrf.go",
        "device.go",
        "logger.go",
        "metadata.go",
        "perf.go",
        "prometheus.go",
        "recovery.go",
        "routergroup.go",
        "server.go",
        "trace.go",
        "utils.go",
    ],
    importpath = "go-common/library/net/http/blademaster",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/conf/dsn",
        "//library/conf/env",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/render",
        "//library/net/ip",
        "//library/net/metadata",
        "//library/net/netutil/breaker",
        "//library/net/trace",
        "//library/stat",
        "//library/time",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/gogo/protobuf/types",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/prometheus/client_golang/prometheus/promhttp",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["example_test.go"],
    tags = ["automanaged"],
    deps = [
        "//library/net/http/blademaster:go_default_library",
        "//library/net/http/blademaster/binding:go_default_library",
        "//library/net/http/blademaster/middleware/auth:go_default_library",
        "//library/net/http/blademaster/middleware/verify:go_default_library",
        "//library/net/http/blademaster/tests:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//library/net/http/blademaster/binding:all-srcs",
        "//library/net/http/blademaster/middleware/antispam:all-srcs",
        "//library/net/http/blademaster/middleware/auth:all-srcs",
        "//library/net/http/blademaster/middleware/cache:all-srcs",
        "//library/net/http/blademaster/middleware/limit/aqm:all-srcs",
        "//library/net/http/blademaster/middleware/permit:all-srcs",
        "//library/net/http/blademaster/middleware/proxy:all-srcs",
        "//library/net/http/blademaster/middleware/rate:all-srcs",
        "//library/net/http/blademaster/middleware/supervisor:all-srcs",
        "//library/net/http/blademaster/middleware/tag:all-srcs",
        "//library/net/http/blademaster/middleware/verify:all-srcs",
        "//library/net/http/blademaster/render:all-srcs",
        "//library/net/http/blademaster/tests:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
