package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "bws",
    srcs = [
        "achievement.go",
        "bws.go",
        "check.go",
        "lottery.go",
        "point.go",
    ],
    importpath = "go-common/app/interface/main/activity/service/bws",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/dao/bws",
        "//app/interface/main/activity/model/bws",
        "//app/service/main/account/api",
        "//app/service/main/usersuit/model",
        "//app/service/main/usersuit/rpc/client",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
        "//library/time",
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
    name = "bws_test",
    srcs = ["bws_test.go"],
    embed = [":bws"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/model/bws",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
