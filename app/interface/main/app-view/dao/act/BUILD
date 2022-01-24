package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "act_test",
    srcs = ["dao_test.go"],
    embed = [":act"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-view/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "act",
    srcs = ["dao.go"],
    importpath = "go-common/app/interface/main/app-view/dao/act",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/activity/model/like",
        "//app/interface/main/activity/rpc/client",
        "//app/interface/main/app-view/conf",
        "//library/ecode",
        "//library/net/http/blademaster",
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
