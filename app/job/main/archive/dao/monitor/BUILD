package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "monitor_test",
    srcs = [
        "dao_test.go",
        "monitor_test.go",
    ],
    embed = [":monitor"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/archive/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "monitor",
    srcs = [
        "dao.go",
        "monitor.go",
    ],
    importpath = "go-common/app/job/main/archive/dao/monitor",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/archive/conf",
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
