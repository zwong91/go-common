package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "cache_test",
    srcs = ["cache_test.go"],
    embed = [":cache"],
    tags = ["automanaged"],
)

go_library(
    name = "cache",
    srcs = ["cache.go"],
    importpath = "go-common/library/cache",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//library/stat/prom",
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
        "//library/cache/memcache:all-srcs",
        "//library/cache/redis:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
