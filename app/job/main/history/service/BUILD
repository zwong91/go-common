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
        "//app/job/main/history/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "delete.go",
        "merge.go",
        "progress.go",
        "service.go",
    ],
    importpath = "go-common/app/job/main/history/service",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/history/model",
        "//app/job/main/history/conf",
        "//app/job/main/history/dao",
        "//app/job/main/history/model",
        "//app/service/main/history/model",
        "//library/log",
        "//library/queue/databus",
        "//library/stat/prom",
        "//library/sync/pipeline",
        "//library/sync/pipeline/fanout",
        "//library/xstr",
        "//vendor/golang.org/x/time/rate",
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
