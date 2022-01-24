package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "like_test.go",
        "service_test.go",
        "stats_test.go",
        "user_likes_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/thumbup/conf",
        "//app/service/main/thumbup/model",
        "//library/cache/redis",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "check.go",
        "has_like.go",
        "item_likes.go",
        "like.go",
        "service.go",
        "stats.go",
        "user_likes.go",
    ],
    importpath = "go-common/app/service/main/thumbup/service",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/thumbup/api",
        "//app/service/main/thumbup/conf",
        "//app/service/main/thumbup/dao",
        "//app/service/main/thumbup/model",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus/report",
        "//library/stat/prom",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
        "//library/time",
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
