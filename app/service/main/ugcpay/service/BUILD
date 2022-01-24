package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "asset.go",
        "income.go",
        "pay.go",
        "service.go",
        "trade.go",
    ],
    importpath = "go-common/app/service/main/ugcpay/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/ugcpay/conf",
        "//app/service/main/ugcpay/dao",
        "//app/service/main/ugcpay/model",
        "//app/service/main/ugcpay/service/pay",
        "//library/cache",
        "//library/database/sql",
        "//library/ecode",
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
    srcs = [
        ":package-srcs",
        "//app/service/main/ugcpay/service/pay:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "service_test",
    srcs = ["income_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/ugcpay/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
