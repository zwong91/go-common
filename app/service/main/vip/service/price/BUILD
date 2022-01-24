package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "price_test",
    srcs = ["price_test.go"],
    embed = [":price"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/vip/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "price",
    srcs = ["price.go"],
    importpath = "go-common/app/service/main/vip/service/price",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//app/service/main/vip/model"],
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
