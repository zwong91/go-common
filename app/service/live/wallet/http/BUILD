package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "http_test",
    srcs = [
        "delCache_test.go",
        "exchange_test.go",
        "getTid_test.go",
        "get_test.go",
        "pay_test.go",
        "recharge_test.go",
    ],
    embed = [":http"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/live/wallet/conf",
        "//app/service/live/wallet/dao",
        "//app/service/live/wallet/model",
        "//app/service/live/wallet/service",
        "//library/ecode",
        "//library/net/http/blademaster",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "http",
    srcs = [
        "http.go",
        "wallet.go",
    ],
    importpath = "go-common/app/service/live/wallet/http",
    tags = ["automanaged"],
    deps = [
        "//app/service/live/wallet/conf",
        "//app/service/live/wallet/model",
        "//app/service/live/wallet/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/verify",
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
