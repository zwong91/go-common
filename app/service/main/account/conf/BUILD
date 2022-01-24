package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "conf",
    srcs = ["conf.go"],
    importpath = "go-common/app/service/main/account/conf",
    tags = ["automanaged"],
    deps = [
        "//library/cache/memcache",
        "//library/conf",
        "//library/database/elastic",
        "//library/log",
        "//library/naming/livezk",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/rpc",
        "//library/net/rpc/warden",
        "//library/net/trace",
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
