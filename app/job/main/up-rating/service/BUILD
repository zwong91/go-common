package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "base.go",
        "paramter.go",
        "past.go",
        "rating.go",
        "run.go",
        "score.go",
        "service.go",
        "statistics.go",
        "task_status.go",
        "top.go",
        "trend.go",
        "trend_heap.go",
    ],
    importpath = "go-common/app/job/main/up-rating/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/up-rating/conf",
        "//app/job/main/up-rating/dao",
        "//app/job/main/up-rating/model",
        "//library/log",
        "//library/time",
        "//vendor/golang.org/x/sync/errgroup",
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
