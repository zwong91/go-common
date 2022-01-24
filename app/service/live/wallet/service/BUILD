package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "delCache_test.go",
        "exchange_test.go",
        "getTid_test.go",
        "get_test.go",
        "local_exchange_test.go",
        "modify_test.go",
        "out_exchange_test.go",
        "pay_test.go",
        "query_test.go",
        "recharge_test.go",
        "recordCoinStream_test.go",
        "service_test.go",
        "wallet_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/live/wallet/conf",
        "//app/service/live/wallet/model",
        "//library/database/sql",
        "//library/ecode",
        "//library/sync/errgroup",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "delCache.go",
        "exchange.go",
        "get.go",
        "getAll.go",
        "getTid.go",
        "local_exchange.go",
        "lock.go",
        "modify.go",
        "out_exchange.go",
        "pay.go",
        "query.go",
        "recharge.go",
        "recordCoinStream.go",
        "service.go",
        "wallet.go",
    ],
    importpath = "go-common/app/service/live/wallet/service",
    tags = ["automanaged"],
    deps = [
        "//app/service/live/wallet/conf",
        "//app/service/live/wallet/dao",
        "//app/service/live/wallet/model",
        "//library/cache",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/ecode",
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
