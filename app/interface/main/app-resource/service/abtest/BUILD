package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "abtest_test",
    srcs = ["abtest_test.go"],
    embed = [":abtest"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-resource/conf",
        "//app/service/main/resource/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "abtest",
    srcs = ["abtest.go"],
    importpath = "go-common/app/interface/main/app-resource/service/abtest",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-resource/conf",
        "//app/interface/main/app-resource/dao/abtest",
        "//app/interface/main/app-resource/model/experiment",
        "//app/service/main/resource/model",
        "//library/log",
        "//vendor/github.com/dgryski/go-farm",
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
