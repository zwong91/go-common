package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "conf",
    srcs = ["conf.go"],
    importpath = "go-common/app/job/main/history/conf",
    tags = ["automanaged"],
    deps = [
        "//library/cache/redis",
        "//library/conf",
        "//library/database/hbase.v2:hbase_v2",
        "//library/database/tidb",
        "//library/ecode/tip",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/trace",
        "//library/queue/databus",
        "//library/sync/pipeline",
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
