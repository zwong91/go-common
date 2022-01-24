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
    data = ["stream-mng.toml"],
    importpath = "go-common/app/service/video/stream-mng/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/video/stream-mng/conf",
        "//app/service/video/stream-mng/server/grpc",
        "//app/service/video/stream-mng/server/http",
        "//app/service/video/stream-mng/service",
        "//library/conf/env",
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
        "//app/service/video/stream-mng/cmd/client:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
