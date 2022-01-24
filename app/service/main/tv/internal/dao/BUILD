package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "account.go",
        "change_history.go",
        "contract.go",
        "dao.cache.go",
        "dao.go",
        "lock.go",
        "main_vip.go",
        "mc.cache.go",
        "mc.cache.key.go",
        "pay_order.go",
        "price_config.go",
        "user_info.go",
        "yst.go",
    ],
    importpath = "go-common/app/service/main/tv/internal/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/account/api",
        "//app/service/main/tv/internal/conf",
        "//app/service/main/tv/internal/model",
        "//app/service/main/tv/internal/pkg",
        "//app/service/main/vipinfo/api",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
        "//library/sync/pipeline/fanout",
        "//library/time",
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

go_test(
    name = "dao_test",
    srcs = [
        "account_test.go",
        "change_history_test.go",
        "contract_test.go",
        "dao.cache_test.go",
        "dao_test.go",
        "lock_test.go",
        "main_vip_test.go",
        "mc.cache.key_test.go",
        "mc.cache_test.go",
        "pay_order_test.go",
        "price_config_test.go",
        "user_info_test.go",
        "yst_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/tv/internal/conf",
        "//app/service/main/tv/internal/model",
        "//library/ecode",
        "//library/net/http/blademaster",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
