package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "conf",
    srcs = ["conf.go"],
    importpath = "go-common/app/interface/openplatform/monitor-end/conf",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/openplatform/monitor-end/model/kafka",
        "//app/interface/openplatform/monitor-end/model/monitor",
        "//library/cache/redis",
        "//library/conf/paladin",
        "//library/database/sql",
        "//library/log",
        "//library/log/infoc",
        "//library/net/http/blademaster",
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
