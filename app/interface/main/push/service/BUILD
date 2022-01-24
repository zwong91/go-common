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
        "//app/interface/main/push/conf",
        "//app/service/main/push/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "callback.go",
        "push.go",
        "report.go",
        "service.go",
        "setting.go",
    ],
    importpath = "go-common/app/interface/main/push/service",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/push/conf",
        "//app/interface/main/push/dao",
        "//app/interface/main/push/model",
        "//app/service/main/push/api/grpc/v1:grpc",
        "//app/service/main/push/dao/huawei",
        "//app/service/main/push/dao/jpush",
        "//app/service/main/push/dao/mi",
        "//app/service/main/push/dao/oppo",
        "//app/service/main/push/model",
        "//library/cache",
        "//library/log",
        "//library/net/http/blademaster",
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
