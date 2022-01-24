package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "wechat_test",
    srcs = ["wechat_test.go"],
    embed = [":wechat"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/ep/saga/conf",
        "//app/admin/ep/saga/dao",
        "//app/admin/ep/saga/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "wechat",
    srcs = [
        "contact.go",
        "wechat.go",
    ],
    importpath = "go-common/app/admin/ep/saga/service/wechat",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/ep/saga/conf",
        "//app/admin/ep/saga/dao",
        "//app/admin/ep/saga/model",
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
