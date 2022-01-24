package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "conf",
    srcs = ["conf.go"],
    importpath = "go-common/app/interface/main/account/conf",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/vip/verify",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/conf",
        "//library/ecode/tip",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/antispam",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/supervisor",
        "//library/net/http/blademaster/middleware/verify",
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
