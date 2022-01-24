package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "aids_test",
    srcs = ["aids_test.go"],
    embed = [":aids"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/app/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "aids",
    srcs = ["aids.go"],
    importpath = "go-common/app/admin/main/app/service/aids",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/app/conf",
        "//app/admin/main/app/dao/aids",
        "//app/admin/main/app/model/aids",
        "//library/ecode",
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
