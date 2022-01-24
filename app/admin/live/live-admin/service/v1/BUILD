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
        "capsule.go",
        "gaea.go",
        "payGoods.go",
        "payLive.go",
        "resource.go",
        "roomMng.go",
        "splash.go",
        "token.go",
        "upload.go",
    ],
    importpath = "go-common/app/admin/live/live-admin/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/live/live-admin/api/http/v1:http",
        "//app/admin/live/live-admin/conf",
        "//app/admin/live/live-admin/dao",
        "//app/admin/live/live-admin/model",
        "//app/service/live/av/api/liverpc/v0:liverpc",
        "//app/service/live/relation/api/liverpc/v1:liverpc",
        "//app/service/live/resource/api/grpc/v1:grpc",
        "//app/service/live/room/api/liverpc",
        "//app/service/live/room/api/liverpc/v1:liverpc",
        "//app/service/live/xlottery/api/grpc/v1:grpc",
        "//app/service/video/stream-mng/api/v1:api",
        "//library/database/bfs",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/sync/errgroup",
        "//vendor/google.golang.org/grpc/status",
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
    srcs = ["token_test.go"],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/live/live-admin/api/http/v1:http",
        "//app/admin/live/live-admin/conf",
        "//app/admin/live/live-admin/dao",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
