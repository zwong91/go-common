package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "permit_test",
    srcs = [
        "example_test.go",
        "permit_test.go",
    ],
    embed = [":permit"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/cache/memcache",
        "//library/container/pool",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/net/netutil/breaker",
        "//library/time",
    ],
)

go_library(
    name = "permit",
    srcs = [
        "permit.go",
        "session.go",
    ],
    importpath = "go-common/library/net/http/blademaster/middleware/permit",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/manager/api",
        "//library/cache/memcache",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/net/rpc/warden",
        "//vendor/github.com/pkg/errors",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["example_test.go"],
    tags = ["automanaged"],
    deps = [
        "//library/cache/memcache:go_default_library",
        "//library/container/pool:go_default_library",
        "//library/net/http/blademaster:go_default_library",
        "//library/net/http/blademaster/middleware/permit:go_default_library",
        "//library/net/metadata:go_default_library",
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
