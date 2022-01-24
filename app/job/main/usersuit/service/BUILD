package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "notify_test.go",
        "pendant_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/usersuit/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "medal.go",
        "notify.go",
        "pendant.go",
        "service.go",
        "vip.go",
    ],
    importpath = "go-common/app/job/main/usersuit/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/usersuit/conf",
        "//app/job/main/usersuit/dao/medal",
        "//app/job/main/usersuit/dao/pendant",
        "//app/job/main/usersuit/model",
        "//app/service/main/vip/model",
        "//library/log",
        "//library/queue/databus",
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
