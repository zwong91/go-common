load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "cmd",
    embed = [":cmd_lib"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "cmd_lib",
    srcs = ["main.go"],
    data = [
        "app_setting.toml",
        "upload_filter.toml",
        "video-c.toml",
    ],
    importpath = "go-common/app/interface/bbq/app-bbq/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/bbq/app-bbq/conf",
        "//app/interface/bbq/app-bbq/server/http",
        "//library/ecode/tip",
        "//library/log",
        "//library/net/trace",
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
