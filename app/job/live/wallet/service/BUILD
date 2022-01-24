package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "service_test.go",
        "wallet_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/live/wallet/conf",
        "//app/job/live/wallet/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "service.go",
        "wallet.go",
    ],
    importpath = "go-common/app/job/live/wallet/service",
    tags = ["automanaged"],
    deps = [
        "//app/job/live/wallet/conf",
        "//app/job/live/wallet/dao",
        "//app/job/live/wallet/model",
        "//library/log",
        "//library/queue/databus",
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
