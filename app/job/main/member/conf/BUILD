package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "conf",
    srcs = ["conf.go"],
    importpath = "go-common/app/job/main/member/conf",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/member/model/block",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/conf",
        "//library/database/sql",
        "//library/log",
        "//library/log/infoc",
        "//library/net/http/blademaster",
        "//library/queue/databus",
        "//library/queue/databus/databusutil",
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
