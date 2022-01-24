package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = ["dao_test.go"],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-player/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = ["dao.go"],
    importpath = "go-common/app/interface/main/app-player/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-player/conf",
        "//app/interface/main/app-player/model",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
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
    srcs = [
        ":package-srcs",
        "//app/interface/main/app-player/dao/account:all-srcs",
        "//app/interface/main/app-player/dao/archive:all-srcs",
        "//app/interface/main/app-player/dao/resource:all-srcs",
        "//app/interface/main/app-player/dao/ugcpay:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
