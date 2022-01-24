package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "banner_test.go",
        "liveCheck_test.go",
        "resource_test.go",
        "splash_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/resource/api/grpc/v1:grpc",
        "//app/service/live/resource/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "banner.go",
        "liveCheck.go",
        "resource.go",
        "splash.go",
        "titans.go",
        "tree.go",
    ],
    importpath = "go-common/app/service/live/resource/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/resource/api/grpc/v1:grpc",
        "//app/service/live/resource/api/http/v1:http",
        "//app/service/live/resource/conf",
        "//app/service/live/resource/dao",
        "//app/service/live/resource/lrucache",
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
