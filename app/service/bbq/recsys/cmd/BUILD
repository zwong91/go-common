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
    data = ["test.toml"],
    importpath = "go-common/app/service/bbq/recsys/cmd",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/recsys/conf",
        "//app/service/bbq/recsys/server/grpc",
        "//app/service/bbq/recsys/server/http",
        "//app/service/bbq/recsys/service",
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
        "//app/service/bbq/recsys/cmd/client:all-srcs",
        "//app/service/bbq/recsys/cmd/relatedrecsys:all-srcs",
        "//app/service/bbq/recsys/cmd/upsrecsys:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
