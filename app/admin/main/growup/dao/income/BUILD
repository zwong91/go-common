package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "income_test",
    srcs = [
        "av_breach_test.go",
        "av_income_test.go",
        "bgm_test.go",
        "blacklist_test.go",
        "charge_test.go",
        "dao_test.go",
        "lottery_test.go",
        "up_account_test.go",
        "up_category_info_test.go",
        "up_income_test.go",
        "up_info_video_test.go",
        "up_withdraw_test.go",
    ],
    embed = [":income"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/growup/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "income",
    srcs = [
        "av_breach.go",
        "av_income.go",
        "bgm.go",
        "blacklist.go",
        "charge.go",
        "dao.go",
        "lottery.go",
        "up_account.go",
        "up_category_info.go",
        "up_income.go",
        "up_info_video.go",
        "up_withdraw.go",
    ],
    importpath = "go-common/app/admin/main/growup/dao/income",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/growup/conf",
        "//app/admin/main/growup/model/income",
        "//library/database/sql",
        "//library/log",
        "//library/time",
        "//library/xstr",
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
