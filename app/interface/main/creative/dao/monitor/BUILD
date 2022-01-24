package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "monitor",
    srcs = ["dao.go"],
    importpath = "go-common/app/interface/main/creative/dao/monitor",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/monitor",
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

go_test(
    name = "monitor_test",
    srcs = ["dao_test.go"],
    embed = [":monitor"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
