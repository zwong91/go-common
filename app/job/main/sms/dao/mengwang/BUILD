package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "mengwang_test",
    srcs = ["client_test.go"],
    embed = [":mengwang"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/sms/conf",
        "//app/service/main/sms/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "mengwang",
    srcs = ["client.go"],
    importpath = "go-common/app/job/main/sms/dao/mengwang",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/sms/conf",
        "//app/job/main/sms/model",
        "//app/service/main/sms/model",
        "//library/log",
        "//library/net/http/blademaster",
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
