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
        "//app/job/main/push/conf",
        "//app/service/main/push/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "callback.go",
        "dataplatform.go",
        "report.go",
        "service.go",
        "task.go",
    ],
    importpath = "go-common/app/job/main/push/service",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/push/model",
        "//app/job/main/push/conf",
        "//app/job/main/push/dao",
        "//app/job/main/push/model",
        "//app/service/main/push/api/grpc/v1:grpc",
        "//app/service/main/push/model",
        "//library/cache",
        "//library/conf/env",
        "//library/database/sql",
        "//library/log",
        "//library/queue/databus",
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
