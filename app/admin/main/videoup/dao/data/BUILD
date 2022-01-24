package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "data_test",
    srcs = ["dao_test.go"],
    embed = [":data"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/videoup/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "data",
    srcs = ["dao.go"],
    importpath = "go-common/app/admin/main/videoup/dao/data",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/videoup/conf",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/xstr",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
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
