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
    data = ["tapd.toml"],
    importpath = "go-common/app/service/ep/footman/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/ep/footman/conf",
        "//app/service/ep/footman/server/http",
        "//app/service/ep/footman/service",
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
    srcs = [
        ":package-srcs",
        "//app/service/ep/footman/cmd/bugly:all-srcs",
        "//app/service/ep/footman/cmd/tapd:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
