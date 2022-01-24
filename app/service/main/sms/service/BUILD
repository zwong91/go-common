package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = ["sms_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/sms/api",
        "//app/service/main/sms/conf",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "service.go",
        "sms.go",
        "task.go",
    ],
    importpath = "go-common/app/service/main/sms/service",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/sms/api",
        "//app/service/main/sms/conf",
        "//app/service/main/sms/dao",
        "//app/service/main/sms/model",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup",
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
