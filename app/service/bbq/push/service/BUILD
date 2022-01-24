package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "message.go",
        "notification.go",
        "parallel.go",
        "service.go",
    ],
    importpath = "go-common/app/service/bbq/push/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/push/api/grpc/v1:grpc",
        "//app/service/bbq/push/conf",
        "//app/service/bbq/push/dao",
        "//app/service/bbq/push/dao/jpush",
        "//app/service/bbq/push/model",
        "//library/log",
        "//library/log/infoc",
        "//library/net/trace",
        "//vendor/github.com/Dai0522/workpool",
        "//vendor/github.com/json-iterator/go",
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
