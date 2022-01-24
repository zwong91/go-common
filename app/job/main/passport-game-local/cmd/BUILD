load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_library(
    name = "cmd_lib",
    srcs = ["main.go"],
    data = ["passport-game-local-job.toml"],
    importpath = "go-common/app/job/main/passport-game-local/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
    deps = [
        "//app/job/main/passport-game-local/conf",
        "//app/job/main/passport-game-local/http",
        "//app/job/main/passport-game-local/service",
        "//library/log",
        "//library/net/trace",
    ],
)

go_binary(
    name = "cmd",
    embed = [":cmd_lib"],
    visibility = ["//visibility:public"],
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
