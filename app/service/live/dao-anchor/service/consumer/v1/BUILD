package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "consumer",
    srcs = [
        "danmu.go",
        "dataBusNotify.go",
        "gift.go",
        "guard.go",
        "popularity.go",
        "rankList.go",
        "roomTag.go",
        "validDays.go",
    ],
    importpath = "go-common/app/service/live/dao-anchor/service/consumer/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/dao-anchor/api/grpc/v1:grpc",
        "//app/service/live/dao-anchor/conf",
        "//app/service/live/dao-anchor/dao",
        "//app/service/live/dao-anchor/model",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/prometheus/common/log",
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
    name = "consumer_test",
    srcs = [
        "popularity_test.go",
        "rankList_test.go",
        "roomTag_test.go",
        "validDays_test.go",
    ],
    embed = [":consumer"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/dao-anchor/api/grpc/v1:grpc",
        "//app/service/live/dao-anchor/conf",
        "//app/service/live/dao-anchor/dao",
        "//library/log",
        "//library/queue/databus",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
