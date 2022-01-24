package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "activity_test",
    srcs = ["dao_test.go"],
    embed = [":activity"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/videoup/conf",
        "//app/job/main/videoup/model/archive",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "activity",
    srcs = ["dao.go"],
    importpath = "go-common/app/job/main/videoup/dao/activity",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/videoup/conf",
        "//app/job/main/videoup/model/archive",
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
