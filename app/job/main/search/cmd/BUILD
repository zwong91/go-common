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
    data = ["search-job-test.toml"],
    importpath = "go-common/app/job/main/search/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/search/conf",
        "//app/job/main/search/http",
        "//app/job/main/search/model",
        "//app/job/main/search/service",
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
