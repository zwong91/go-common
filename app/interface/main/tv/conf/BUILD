load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "conf",
    srcs = [
        "cache.go",
        "conf.go",
        "homepage.go",
        "search_sug.go",
        "zone.go",
    ],
    importpath = "go-common/app/interface/main/tv/conf",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/tv/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/conf",
        "//library/database/sql",
        "//library/ecode/tip",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
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
