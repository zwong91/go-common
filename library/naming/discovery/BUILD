package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "discovery_test",
    srcs = [
        "discovery_test.go",
        "example_test.go",
    ],
    embed = [":discovery"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/conf/env",
        "//library/ecode",
        "//library/exp/feature",
        "//library/naming",
        "//library/net/http/blademaster",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "discovery",
    srcs = [
        "discovery.go",
        "shuffle.go",
    ],
    importpath = "go-common/library/naming/discovery",
    tags = ["automanaged"],
    deps = [
        "//library/conf/env",
        "//library/ecode",
        "//library/exp/feature",
        "//library/log",
        "//library/naming",
        "//library/net/http/blademaster",
        "//library/net/netutil",
        "//library/net/netutil/breaker",
        "//library/time",
        "//library/xstr",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["example_test.go"],
    tags = ["automanaged"],
    deps = [
        "//library/naming:go_default_library",
        "//library/naming/discovery:go_default_library",
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
