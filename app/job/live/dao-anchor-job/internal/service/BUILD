package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "cover_test.go",
        "data_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/job/live/dao-anchor-job/internal/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "cover.go",
        "data.go",
        "service.go",
    ],
    importpath = "go-common/app/job/live/dao-anchor-job/internal/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/live/dao-anchor-job/internal/conf",
        "//app/job/live/dao-anchor-job/internal/dao",
        "//library/log",
        "//library/net/trace",
        "//library/sync/errgroup",
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
