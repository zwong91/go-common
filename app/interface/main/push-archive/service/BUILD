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
        "limit_test.go",
        "service_test.go",
        "setting_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/push-archive/conf",
        "//app/interface/main/push-archive/dao",
        "//app/interface/main/push-archive/model",
        "//app/service/main/push/api/grpc/v1:grpc",
        "//library/ecode",
        "//library/net/http/blademaster",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "abtest.go",
        "archive.go",
        "limit.go",
        "relation.go",
        "service.go",
        "setting.go",
        "statistics.go",
    ],
    importpath = "go-common/app/interface/main/push-archive/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/push-archive/conf",
        "//app/interface/main/push-archive/dao",
        "//app/interface/main/push-archive/model",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//app/service/main/push/api/grpc/v1:grpc",
        "//library/cache",
        "//library/conf/env",
        "//library/log",
        "//library/net/http/blademaster",
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
