package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "vip_test",
    srcs = ["vip_test.go"],
    embed = [":vip"],
    tags = ["automanaged"],
)

go_library(
    name = "vip",
    srcs = ["vip.go"],
    importpath = "go-common/app/service/live/xuser/service/vip/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/xuser/api/grpc/v1:grpc",
        "//app/service/live/xuser/conf",
        "//app/service/live/xuser/dao/vip",
        "//app/service/live/xuser/model",
        "//library/cache",
        "//library/log",
        "//library/net/metadata",
        "//library/queue/databus",
        "//library/time",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/pkg/errors",
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
