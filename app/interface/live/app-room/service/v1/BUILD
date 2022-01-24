load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "gift_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/live/app-room/api/http/v1:http",
        "//app/interface/live/app-room/conf",
        "//app/interface/live/app-room/model",
        "//library/log",
        "//library/net/metadata",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "banner.go",
        "gift.go",
        "roomNotice.go",
    ],
    importpath = "go-common/app/interface/live/app-room/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/app-room/api/http/v1:http",
        "//app/interface/live/app-room/conf",
        "//app/interface/live/app-room/dao",
        "//app/interface/live/app-room/model",
        "//app/service/live/gift/api/grpc/v1:grpc",
        "//app/service/live/resource/api/grpc/v1:grpc",
        "//app/service/live/xuserex/api/grpc/v1:grpc",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
    srcs = [
        ":package-srcs",
        "//app/interface/live/app-room/service/v1/dm:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
