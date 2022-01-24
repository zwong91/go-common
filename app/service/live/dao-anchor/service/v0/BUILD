package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = ["createData_test.go"],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/dao-anchor/api/grpc/v0:grpc",
        "//app/service/live/dao-anchor/conf",
        "//app/service/live/dao-anchor/dao",
        "//library/net/metadata",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = ["createData.go"],
    importpath = "go-common/app/service/live/dao-anchor/service/v0",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/dao-anchor/api/grpc/v0:grpc",
        "//app/service/live/dao-anchor/api/grpc/v1:grpc",
        "//app/service/live/dao-anchor/conf",
        "//app/service/live/dao-anchor/dao",
        "//library/cache/redis",
        "//library/log",
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
