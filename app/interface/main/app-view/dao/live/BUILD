package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "live_test",
    srcs = ["dao_test.go"],
    embed = [":live"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-view/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "live",
    srcs = ["dao.go"],
    importpath = "go-common/app/interface/main/app-view/dao/live",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-view/conf",
        "//app/interface/main/app-view/model/live",
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
