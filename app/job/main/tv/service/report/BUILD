package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "report_test",
    srcs = ["service_test.go"],
    embed = [":report"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//app/job/main/tv/conf"],
)

go_library(
    name = "report",
    srcs = [
        "service.go",
        "style.go",
    ],
    importpath = "go-common/app/job/main/tv/service/report",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/job/main/tv/dao/report",
        "//app/job/main/tv/model/pgc",
        "//app/job/main/tv/model/report",
        "//library/log",
        "//library/sync/pipeline/fanout",
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
