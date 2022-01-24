package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dm_test",
    srcs = [
        "broadcast_test.go",
        "dao_test.go",
    ],
    embed = [":dm"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/activity/conf",
        "//app/job/main/activity/model/dm",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dm",
    srcs = [
        "broadcast.go",
        "dao.go",
    ],
    importpath = "go-common/app/job/main/activity/dao/dm",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/activity/conf",
        "//app/job/main/activity/model/dm",
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
