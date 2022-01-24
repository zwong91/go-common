package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "bag_test.go",
        "databus_test.go",
        "discount_test.go",
        "gift_test.go",
        "infoc_test.go",
        "list_test.go",
        "package_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/gift/api/grpc/v1:grpc",
        "//app/service/live/gift/conf",
        "//app/service/live/gift/dao",
        "//app/service/live/gift/model",
        "//app/service/live/resource/sdk",
        "//vendor/github.com/satori/go.uuid:go_uuid",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "bag.go",
        "databus.go",
        "discount.go",
        "gift.go",
        "infoc.go",
        "list.go",
        "package.go",
    ],
    importpath = "go-common/app/service/live/gift/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/activity/api/liverpc/v1:liverpc",
        "//app/service/live/fans_medal/api/liverpc/v2:liverpc",
        "//app/service/live/gift/api/grpc/v1:grpc",
        "//app/service/live/gift/conf",
        "//app/service/live/gift/dao",
        "//app/service/live/gift/model",
        "//app/service/live/live_user/api/liverpc/v1:liverpc",
        "//app/service/live/resource/sdk",
        "//app/service/live/room/api/liverpc/v1:liverpc",
        "//app/service/live/user/api/liverpc/v3:liverpc",
        "//app/service/live/xuser/api/grpc/v1:grpc",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//library/net/rpc/liverpc/context",
        "//library/queue/databus",
        "//library/sync/errgroup",
        "//vendor/github.com/satori/go.uuid:go_uuid",
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
