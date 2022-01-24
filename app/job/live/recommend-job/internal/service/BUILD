package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "item_cf_job.go",
        "service.go",
        "user_area_job.go",
    ],
    importpath = "go-common/app/job/live/recommend-job/internal/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/live/recommend-job/internal/conf",
        "//app/job/live/recommend-job/internal/dao",
        "//app/service/live/recommend/recconst",
        "//library/cache/redis",
        "//library/log",
        "//library/sync/errgroup",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/robfig/cron",
        "//vendor/github.com/siddontang/go/ioutil2",
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
