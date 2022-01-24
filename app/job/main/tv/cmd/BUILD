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
    data = ["tv-job-test.toml"],
    importpath = "go-common/app/job/main/tv/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/job/main/tv/http",
        "//app/job/main/tv/service/pgc",
        "//app/job/main/tv/service/ugc",
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
