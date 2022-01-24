package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "ping_test",
    srcs = ["ping_test.go"],
    embed = [":ping"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-feed/conf",
        "//app/interface/main/app-feed/dao/archive",
        "//app/interface/main/app-feed/dao/audit",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "ping",
    srcs = ["ping.go"],
    importpath = "go-common/app/interface/main/app-feed/service/ping",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-feed/conf",
        "//app/interface/main/app-feed/dao/archive",
        "//app/interface/main/app-feed/dao/audit",
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
