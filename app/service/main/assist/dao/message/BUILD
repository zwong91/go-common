package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "message_test",
    srcs = ["dao_test.go"],
    embed = [":message"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//app/service/main/assist/conf"],
)

go_library(
    name = "message",
    srcs = ["dao.go"],
    importpath = "go-common/app/service/main/assist/dao/message",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/assist/conf",
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
