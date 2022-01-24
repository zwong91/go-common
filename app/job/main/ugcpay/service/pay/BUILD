package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "pay_test",
    srcs = ["pay_test.go"],
    embed = [":pay"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/ugcpay/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "pay",
    srcs = ["pay.go"],
    importpath = "go-common/app/job/main/ugcpay/service/pay",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//vendor/github.com/pkg/errors",
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
