package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "bloomfilter.go",
        "forward_index.go",
        "inverted_index.go",
        "service.go",
    ],
    importpath = "go-common/app/job/bbq/recall/internal/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/bbq/recall/internal/conf",
        "//app/job/bbq/recall/internal/dao",
        "//app/job/bbq/recall/proto",
        "//library/log",
        "//vendor/github.com/golang/snappy",
        "//vendor/github.com/robfig/cron",
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
