package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "zlimit_test",
    srcs = ["zlimit_test.go"],
    embed = [":zlimit"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/cache/redis",
        "//library/container/pool",
        "//library/database/sql",
        "//library/time",
    ],
)

go_library(
    name = "zlimit",
    srcs = [
        "cache.go",
        "db.go",
        "zlimit.go",
    ],
    importpath = "go-common/app/service/main/location/zlimit",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/location/model",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/ip",
        "//library/time",
        "//library/xstr",
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
