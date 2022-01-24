package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "conf_test",
    srcs = ["conf_test.go"],
    embed = [":conf"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "conf",
    srcs = ["conf.go"],
    importpath = "go-common/app/service/main/spy/conf",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/conf",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/rpc",
        "//library/net/rpc/warden",
        "//library/net/trace",
        "//library/queue/databus",
        "//library/time",
        "//vendor/github.com/BurntSushi/toml",
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
