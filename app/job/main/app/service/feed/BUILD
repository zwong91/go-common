package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "feed_test",
    srcs = ["feed_test.go"],
    embed = [":feed"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/app/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "feed",
    srcs = ["feed.go"],
    importpath = "go-common/app/job/main/app/service/feed",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/app/conf",
        "//app/job/main/app/dao/feed",
        "//app/job/main/app/model/feed",
        "//app/service/main/archive/api",
        "//library/log",
        "//library/sync/errgroup",
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
