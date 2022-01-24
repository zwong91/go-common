package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "banner.go",
        "room.go",
        "splash.go",
        "topic.go",
    ],
    importpath = "go-common/app/interface/live/app-blink/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/app-blink/api/http/v1:http",
        "//app/interface/live/app-blink/conf",
        "//app/interface/live/app-blink/dao",
        "//app/service/live/resource/api/grpc/v1:grpc",
        "//library/ecode",
        "//library/net/metadata",
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

go_test(
    name = "service_test",
    srcs = [
        "banner_test.go",
        "room_test.go",
        "splash_test.go",
        "topic_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/live/app-blink/api/http/v1:http",
        "//app/interface/live/app-blink/conf",
        "//app/service/live/resource/api/grpc/v1:grpc",
        "//library/net/metadata",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
