package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "live_test",
    srcs = ["live_test.go"],
    embed = [":live"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-channel/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "live",
    srcs = ["live.go"],
    importpath = "go-common/app/interface/main/app-channel/dao/live",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-card/model/card/live",
        "//app/interface/main/app-channel/conf",
        "//app/interface/main/app-channel/model/live",
        "//library/ecode",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/xstr",
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
