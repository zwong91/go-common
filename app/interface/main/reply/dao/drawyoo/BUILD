package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "drawyoo_test",
    srcs = ["drawyoo_test.go"],
    embed = [":drawyoo"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/reply/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "drawyoo",
    srcs = ["drawyoo.go"],
    importpath = "go-common/app/interface/main/reply/dao/drawyoo",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/reply/conf",
        "//app/interface/main/reply/model/drawyoo",
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
