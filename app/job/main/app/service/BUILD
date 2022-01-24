package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "archive_test.go",
        "contribute_test.go",
        "monitor_test.go",
        "retry_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/app/conf",
        "//app/job/main/app/model/space",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "account.go",
        "archive.go",
        "contribute.go",
        "flush.go",
        "monitor.go",
        "retry.go",
        "service.go",
        "stat.go",
    ],
    importpath = "go-common/app/job/main/app/service",
    tags = ["automanaged"],
    deps = [
        "//app/interface/openplatform/article/model",
        "//app/job/main/app/conf",
        "//app/job/main/app/dao/monitor",
        "//app/job/main/app/dao/space",
        "//app/job/main/app/dao/view",
        "//app/job/main/app/model",
        "//app/job/main/app/model/space",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/resource/model",
        "//app/service/main/resource/rpc/client",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//library/sync/errgroup",
        "//library/time",
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
    srcs = [
        ":package-srcs",
        "//app/job/main/app/service/feed:all-srcs",
        "//app/job/main/app/service/show:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
