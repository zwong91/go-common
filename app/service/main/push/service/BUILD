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
        "//app/service/main/push/conf",
        "//app/service/main/push/model",
        "//library/cache/redis",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "business.go",
        "callback.go",
        "progress.go",
        "push.go",
        "report.go",
        "service.go",
        "setting.go",
        "task.go",
        "upload.go",
    ],
    importpath = "go-common/app/service/main/push/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/filter/model/rpc",
        "//app/service/main/filter/rpc/client",
        "//app/service/main/push/conf",
        "//app/service/main/push/dao",
        "//app/service/main/push/dao/apns2",
        "//app/service/main/push/dao/fcm",
        "//app/service/main/push/dao/huawei",
        "//app/service/main/push/dao/jpush",
        "//app/service/main/push/dao/oppo",
        "//app/service/main/push/model",
        "//library/cache",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/sync/errgroup",
        "//library/time",
        "//library/xstr",
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
