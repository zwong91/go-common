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
    deps = ["//vendor/github.com/smartystreets/goconvey/convey"],
)

go_library(
    name = "service",
    srcs = [
        "achive.go",
        "service.go",
    ],
    importpath = "go-common/app/job/main/feed/service",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/feed/conf",
        "//app/job/main/feed/dao",
        "//app/job/main/feed/model",
        "//app/service/main/feed/model",
        "//app/service/main/feed/rpc/client",
        "//library/conf/env",
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
