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
        "//app/job/main/block/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "audit_log.go",
        "blackhouse.go",
        "limit.go",
        "msg.go",
        "notify.go",
        "service.go",
    ],
    importpath = "go-common/app/job/main/block/service",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/block/conf",
        "//app/job/main/block/dao",
        "//app/job/main/block/model",
        "//library/cache",
        "//library/database/sql",
        "//library/log",
        "//library/queue/databus",
        "//library/queue/databus/report",
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
