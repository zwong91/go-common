package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "common_message_test.go",
        "mids_test.go",
        "push_test.go",
        "service_test.go",
        "start_live_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/live/push-live/conf",
        "//app/interface/live/push-live/dao",
        "//app/interface/live/push-live/model",
        "//library/cache/redis",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "common_message.go",
        "mids.go",
        "push.go",
        "service.go",
        "start_live.go",
    ],
    importpath = "go-common/app/interface/live/push-live/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/push-live/conf",
        "//app/interface/live/push-live/dao",
        "//app/interface/live/push-live/model",
        "//library/cache/redis",
        "//library/log",
        "//library/queue/databus",
        "//library/sync/errgroup",
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
