package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "workflow",
    srcs = ["workflow.go"],
    importpath = "go-common/app/interface/main/reply/dao/workflow",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/reply/conf",
        "//app/interface/main/reply/model/reply",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/xstr",
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
    name = "workflow_test",
    srcs = ["workflow_test.go"],
    embed = [":workflow"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/reply/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
