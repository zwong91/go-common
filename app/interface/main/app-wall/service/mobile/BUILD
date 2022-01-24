package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "mobile_test",
    srcs = ["mobile_test.go"],
    embed = [":mobile"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-wall/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "mobile",
    srcs = [
        "mobile.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/app-wall/service/mobile",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-wall/conf",
        "//app/interface/main/app-wall/dao/mobile",
        "//app/interface/main/app-wall/model",
        "//app/interface/main/app-wall/model/mobile",
        "//library/ecode",
        "//library/log",
        "//library/stat/prom",
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
