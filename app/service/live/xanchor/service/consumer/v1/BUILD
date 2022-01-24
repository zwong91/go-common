package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "consumer",
    srcs = [
        "ConsumerDanmu.go",
        "ConsumerGift.go",
        "dataBusNotify.go",
    ],
    importpath = "go-common/app/service/live/xanchor/service/consumer/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/xanchor/conf",
        "//app/service/live/xanchor/dao",
        "//app/service/live/xanchor/dao/consumer",
        "//app/service/live/xanchor/model",
        "//library/log",
        "//library/queue/databus",
        "//vendor/github.com/prometheus/common/log",
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
