package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "conf",
    srcs = ["conf.go"],
    importpath = "go-common/app/job/main/mcn/conf",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/cache/memcache",
        "//library/conf",
        "//library/database/sql",
        "//library/ecode/tip",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/rpc/warden",
        "//library/net/trace",
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
