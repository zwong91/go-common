package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = ["service_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/click/conf",
        "//app/job/main/click/model",
        "//library/sync/errgroup",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "click.go",
        "forbid.go",
        "redis.go",
        "service.go",
    ],
    importpath = "go-common/app/job/main/click/service",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/click/conf",
        "//app/job/main/click/dao",
        "//app/job/main/click/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//library/cache/redis",
        "//library/log",
        "//library/log/infoc",
        "//library/queue/databus",
        "//vendor/github.com/dgryski/go-farm",
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
