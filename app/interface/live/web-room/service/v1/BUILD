package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "captcha.go",
        "roomAdmin.go",
    ],
    importpath = "go-common/app/interface/live/web-room/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/web-room/api/http/v1:http",
        "//app/interface/live/web-room/conf",
        "//app/interface/live/web-room/dao",
        "//app/service/live/xcaptcha/api/grpc/v1:grpc",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
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
    srcs = [
        ":package-srcs",
        "//app/interface/live/web-room/service/v1/dm:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "service_test",
    srcs = ["roomAdmin_test.go"],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/live/web-room/api/http/v1:http",
        "//app/interface/live/web-room/conf",
        "//library/net/metadata",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
