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
        "//app/job/main/archive/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "archive.go",
        "archive_result.go",
        "cache.go",
        "dm.go",
        "retry.go",
        "service.go",
        "video.go",
    ],
    importpath = "go-common/app/job/main/archive/service",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/dm2/model",
        "//app/interface/main/dm2/rpc/client",
        "//app/job/main/archive/conf",
        "//app/job/main/archive/dao/archive",
        "//app/job/main/archive/dao/email",
        "//app/job/main/archive/dao/monitor",
        "//app/job/main/archive/dao/reply",
        "//app/job/main/archive/dao/result",
        "//app/job/main/archive/model/archive",
        "//app/job/main/archive/model/databus",
        "//app/job/main/archive/model/dm",
        "//app/job/main/archive/model/result",
        "//app/job/main/archive/model/retry",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
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
