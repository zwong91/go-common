package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "income_test.go",
        "notice_test.go",
        "service_test.go",
        "withdraw_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/growup/conf",
        "//app/interface/main/growup/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "activity.go",
        "archive_stat.go",
        "banner.go",
        "bgm.go",
        "column.go",
        "exchange.go",
        "income.go",
        "notice.go",
        "service.go",
        "snow_flake.go",
        "special_award.go",
        "up_bill.go",
        "up_status.go",
        "up_year.go",
        "withdraw.go",
    ],
    importpath = "go-common/app/interface/main/growup/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/growup/conf",
        "//app/interface/main/growup/dao",
        "//app/interface/main/growup/model",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/time",
        "//library/xstr",
        "//vendor/golang.org/x/sync/errgroup",
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
        "//app/interface/main/growup/service/newbie:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
