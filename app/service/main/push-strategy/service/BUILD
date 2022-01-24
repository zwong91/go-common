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
        "//app/service/main/push-strategy/conf",
        "//app/service/main/push/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "service.go",
        "task.go",
    ],
    importpath = "go-common/app/service/main/push-strategy/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/filter/model/rpc",
        "//app/service/main/filter/rpc/client",
        "//app/service/main/push-strategy/conf",
        "//app/service/main/push-strategy/dao",
        "//app/service/main/push-strategy/model",
        "//app/service/main/push/model",
        "//library/cache",
        "//library/ecode",
        "//library/log",
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
