package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "conf",
    srcs = [
        "canal_conf.go",
        "conf.go",
        "tidb_conf.go",
    ],
    importpath = "go-common/app/infra/canal/conf",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/infra/canal/infoc",
        "//library/conf",
        "//library/database/sql",
        "//library/log",
        "//library/naming/discovery",
        "//library/net/http/blademaster",
        "//library/queue/databus",
        "//library/time",
        "//vendor/github.com/BurntSushi/toml",
        "//vendor/github.com/siddontang/go-mysql/canal",
        "//vendor/github.com/siddontang/go-mysql/client",
        "//vendor/github.com/siddontang/go-mysql/mysql",
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
