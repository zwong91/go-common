package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "account_test.go",
        "coin_test.go",
        "reply_test.go",
        "service_test.go",
        "spy_test.go",
        "vip_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/figure/conf",
        "//app/job/main/figure/model",
        "//app/job/main/reply/model/reply",
        "//app/service/main/coin/model",
        "//app/service/main/spy/model",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "account.go",
        "blocked.go",
        "coin.go",
        "danmaku.go",
        "pay.go",
        "reply.go",
        "service.go",
        "spy.go",
        "sync.go",
        "vip.go",
    ],
    importpath = "go-common/app/job/main/figure/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/figure/conf",
        "//app/job/main/figure/dao",
        "//app/job/main/figure/model",
        "//app/service/main/coin/model",
        "//app/service/main/spy/model",
        "//library/log",
        "//library/queue/databus",
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
