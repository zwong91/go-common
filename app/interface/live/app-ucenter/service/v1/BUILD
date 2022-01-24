package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "roomAdmin_test.go",
        "room_test.go",
        "topic_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/live/app-ucenter/api/http/v1:http",
        "//app/interface/live/app-ucenter/conf",
        "//library/net/metadata",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "room.go",
        "roomAdmin.go",
        "topic.go",
    ],
    importpath = "go-common/app/interface/live/app-ucenter/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/app-ucenter/api/http/v1:http",
        "//app/interface/live/app-ucenter/conf",
        "//app/interface/live/app-ucenter/dao",
        "//app/service/live/xuser/api/grpc/v1:grpc",
        "//library/ecode",
        "//library/log",
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
