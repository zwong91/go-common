load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

package(default_visibility = ["//visibility:public"])

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

go_library(
    name = "conf",
    srcs = ["conf.go"],
    importpath = "go-common/app/service/main/msm/conf",
    tags = ["automanaged"],
    deps = [
        "//library/conf",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/rpc",
        "//vendor/github.com/BurntSushi/toml",
    ],
)
