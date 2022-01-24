package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "wechat",
    srcs = [
        "contact.go",
        "wechat.go",
    ],
    importpath = "go-common/app/tool/saga/service/wechat",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/saga/conf",
        "//app/tool/saga/dao",
        "//app/tool/saga/model",
        "//library/log",
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

go_test(
    name = "wechat_test",
    srcs = ["wechat_test.go"],
    embed = [":wechat"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/tool/saga/conf",
        "//app/tool/saga/dao",
        "//app/tool/saga/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
