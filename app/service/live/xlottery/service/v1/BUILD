package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "capsule.go",
        "storm.go",
    ],
    importpath = "go-common/app/service/live/xlottery/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/banned_service/api/liverpc",
        "//app/service/live/banned_service/api/liverpc/v1:liverpc",
        "//app/service/live/broadcast-proxy/api/v1:api",
        "//app/service/live/captcha/api/liverpc",
        "//app/service/live/captcha/api/liverpc/v0:liverpc",
        "//app/service/live/gift/api/liverpc",
        "//app/service/live/gift/api/liverpc/v1:liverpc",
        "//app/service/live/live-dm/api/grpc/v1:grpc",
        "//app/service/live/rc/api/liverpc/v1:liverpc",
        "//app/service/live/userext/api/liverpc/v1:liverpc",
        "//app/service/live/xlottery/api/grpc/v1:grpc",
        "//app/service/live/xlottery/conf",
        "//app/service/live/xlottery/dao",
        "//app/service/live/xlottery/model",
        "//app/service/live/xuser/api/grpc/v1:grpc",
        "//app/service/main/account/api",
        "//app/service/main/account/model",
        "//app/service/main/filter/api/grpc/v1:grpc",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/queue/databus/report",
        "//library/sync/pipeline/fanout",
        "//vendor/github.com/pkg/errors",
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
    srcs = ["capsule_test.go"],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/xlottery/api/grpc/v1:grpc",
        "//app/service/live/xlottery/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
