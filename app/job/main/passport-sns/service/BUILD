package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "service.go",
        "sns_log.go",
        "sync_sns.go",
    ],
    importpath = "go-common/app/job/main/passport-sns/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/passport-sns/conf",
        "//app/job/main/passport-sns/dao",
        "//app/job/main/passport-sns/model",
        "//library/log",
        "//library/queue/databus",
        "//library/queue/databus/databusutil",
        "//library/sync/pipeline/fanout",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/pkg/errors",
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
