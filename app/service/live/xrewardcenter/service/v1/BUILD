package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = ["anchorReward.go"],
    importpath = "go-common/app/service/live/xrewardcenter/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/gift/api/liverpc/v0:liverpc",
        "//app/service/live/room/api/liverpc/v2:liverpc",
        "//app/service/live/xrewardcenter/api/grpc/v1:grpc",
        "//app/service/live/xrewardcenter/conf",
        "//app/service/live/xrewardcenter/dao",
        "//app/service/live/xrewardcenter/dao/anchorReward",
        "//app/service/live/xrewardcenter/model/anchorTask",
        "//library/ecode",
        "//library/log",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/robfig/cron",
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
    srcs = ["anchorReward_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/live/xrewardcenter/api/grpc/v1:grpc",
        "//app/service/live/xrewardcenter/conf",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
