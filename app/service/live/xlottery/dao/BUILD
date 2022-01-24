package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "cache_test.go",
        "storm_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/xlottery/model",
        "//library/cache/redis",
        "//library/container/pool",
        "//library/database/sql",
        "//library/time",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "api.go",
        "award.go",
        "cache.go",
        "capsule.go",
        "capsule_extra.go",
        "coin.go",
        "coin_config.go",
        "dao.go",
        "pool.go",
        "pool_prize.go",
        "prize_extra.go",
        "prize_log.go",
        "storm.go",
    ],
    importpath = "go-common/app/service/live/xlottery/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/gift/api/liverpc",
        "//app/service/live/rc/api/liverpc",
        "//app/service/live/rc/api/liverpc/v1:liverpc",
        "//app/service/live/userext/api/liverpc",
        "//app/service/live/xlottery/api/grpc/v1:grpc",
        "//app/service/live/xlottery/conf",
        "//app/service/live/xlottery/model",
        "//app/service/main/account/rpc/client",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/net/rpc/liverpc",
        "//library/queue/databus/report",
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
