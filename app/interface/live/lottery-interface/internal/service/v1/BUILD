package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = ["storm.go"],
    importpath = "go-common/app/interface/live/lottery-interface/internal/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/lottery-interface/internal/service",
        "//app/service/live/xlottery/api/grpc/v1:grpc",
        "//library/log",
        "//library/log/infoc",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/render",
        "//library/net/metadata",
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
