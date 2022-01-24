package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "splash_test",
    srcs = ["splash_test.go"],
    embed = [":splash"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-resource/conf",
        "//app/interface/main/app-resource/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "splash",
    srcs = ["splash.go"],
    importpath = "go-common/app/interface/main/app-resource/service/splash",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-resource/conf",
        "//app/interface/main/app-resource/dao/ad",
        "//app/interface/main/app-resource/dao/location",
        "//app/interface/main/app-resource/dao/splash",
        "//app/interface/main/app-resource/model",
        "//app/interface/main/app-resource/model/splash",
        "//app/service/main/location/model",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup",
        "//vendor/github.com/dgryski/go-farm",
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
