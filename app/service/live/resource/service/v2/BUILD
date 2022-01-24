package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = ["userResource.go"],
    importpath = "go-common/app/service/live/resource/service/v2",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/resource/api/grpc/v2:grpc",
        "//app/service/live/resource/conf",
        "//app/service/live/resource/dao",
        "//app/service/live/resource/model",
        "//library/ecode",
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

go_test(
    name = "service_test",
    srcs = ["userResource_test.go"],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/resource/api/grpc/v2:grpc",
        "//app/service/live/resource/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
