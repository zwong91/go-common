package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "cmd",
    embed = [":cmd_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "cmd_lib",
    srcs = ["main.go"],
    data = [
        "app-resource-example.toml",
        "app-resource-test.toml",
    ],
    importpath = "go-common/app/interface/main/app-resource/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-resource/conf",
        "//app/interface/main/app-resource/grpc",
        "//app/interface/main/app-resource/http",
        "//app/interface/main/app-resource/service/abtest",
        "//app/interface/main/app-resource/service/audit",
        "//app/interface/main/app-resource/service/broadcast",
        "//app/interface/main/app-resource/service/domain",
        "//app/interface/main/app-resource/service/guide",
        "//app/interface/main/app-resource/service/module",
        "//app/interface/main/app-resource/service/notice",
        "//app/interface/main/app-resource/service/param",
        "//app/interface/main/app-resource/service/ping",
        "//app/interface/main/app-resource/service/plugin",
        "//app/interface/main/app-resource/service/show",
        "//app/interface/main/app-resource/service/sidebar",
        "//app/interface/main/app-resource/service/splash",
        "//app/interface/main/app-resource/service/static",
        "//app/interface/main/app-resource/service/version",
        "//app/interface/main/app-resource/service/white",
        "//library/ecode/tip",
        "//library/log",
        "//library/net/http/blademaster/middleware/auth",
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
