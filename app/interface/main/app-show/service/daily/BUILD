package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "daily_test",
    srcs = ["daily_test.go"],
    embed = [":daily"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-show/conf",
        "//app/interface/main/app-show/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "daily",
    srcs = ["daily.go"],
    importpath = "go-common/app/interface/main/app-show/service/daily",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-show/conf",
        "//app/interface/main/app-show/dao/archive",
        "//app/interface/main/app-show/dao/card",
        "//app/interface/main/app-show/dao/tag",
        "//app/interface/main/app-show/model",
        "//app/interface/main/app-show/model/card",
        "//app/interface/main/app-show/model/daily",
        "//app/service/main/archive/api",
        "//library/log",
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
