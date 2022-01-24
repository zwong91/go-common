package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "conf",
    srcs = ["conf.go"],
    importpath = "go-common/app/admin/main/config/conf",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/conf",
        "//library/database/orm",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/permit",
        "//library/net/rpc",
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
