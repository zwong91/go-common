package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "coinexchange_test.go",
        "coinstream_test.go",
        "mc_wallet_test.go",
        "metal_test.go",
        "pub_test.go",
        "redis_test.go",
        "tx_test.go",
        "wallet_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/live/wallet/conf",
        "//app/service/live/wallet/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/ecode",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "coinexchange.go",
        "coinstream.go",
        "dao.go",
        "mc_wallet.go",
        "metal.go",
        "pub.go",
        "redis.go",
        "tx.go",
        "wallet.go",
    ],
    importpath = "go-common/app/service/live/wallet/dao",
    tags = ["automanaged"],
    deps = [
        "//app/service/live/wallet/conf",
        "//app/service/live/wallet/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/queue/databus",
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
