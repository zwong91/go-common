package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "account_test.go",
        "activity_test.go",
        "archive_stat_test.go",
        "av_breach_test.go",
        "banner_test.go",
        "bgm_test.go",
        "column_income_test.go",
        "dao_test.go",
        "exchange_test.go",
        "hbase_test.go",
        "income_test.go",
        "notice_test.go",
        "redis_test.go",
        "special_award_test.go",
        "tool_test.go",
        "up_bill_test.go",
        "up_test.go",
        "withdraw_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/growup/conf",
        "//app/interface/main/growup/model",
        "//library/net/metadata",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "account.go",
        "activity.go",
        "archive_stat.go",
        "av_breach.go",
        "banner.go",
        "bgm.go",
        "column_income.go",
        "dao.go",
        "exchange.go",
        "hbase.go",
        "income.go",
        "notice.go",
        "redis.go",
        "special_award.go",
        "tool.go",
        "up.go",
        "up_bill.go",
        "up_year.go",
        "withdraw.go",
    ],
    importpath = "go-common/app/interface/main/growup/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/growup/conf",
        "//app/interface/main/growup/model",
        "//app/interface/openplatform/article/model",
        "//app/interface/openplatform/article/rpc/client",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//app/service/main/vip/model",
        "//app/service/main/vip/rpc/client",
        "//library/cache/redis",
        "//library/database/hbase.v2:hbase_v2",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/tsuna/gohbase/hrpc",
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
        "//app/interface/main/growup/dao/newbiedao:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
