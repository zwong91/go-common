package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "notice_test",
    srcs = ["notice_test.go"],
    embed = [":notice"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-resource/conf",
        "//app/interface/main/app-resource/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "notice",
    srcs = ["notice.go"],
    importpath = "go-common/app/interface/main/app-resource/service/notice",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-resource/conf",
        "//app/interface/main/app-resource/dao/location",
        "//app/interface/main/app-resource/dao/notice",
        "//app/interface/main/app-resource/model",
        "//app/interface/main/app-resource/model/notice",
        "//app/service/main/location/model",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
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
