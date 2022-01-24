package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "jpush_test",
    srcs = ["client_test.go"],
    embed = [":jpush"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/smartystreets/goconvey/convey"],
)

go_library(
    name = "jpush",
    srcs = [
        "audience.go",
        "callback.go",
        "client.go",
        "errcode.go",
        "message.go",
        "notice.go",
        "option.go",
        "payload.go",
        "platform.go",
        "report.go",
        "schedule.go",
    ],
    importpath = "go-common/app/service/main/push/dao/jpush",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/stat",
        "//library/stat/prom",
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
