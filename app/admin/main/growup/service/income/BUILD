package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "income_test",
    srcs = [
        "archive_income_test.go",
        "av_breach_test.go",
        "blacklist_test.go",
        "service_test.go",
        "up_account_test.go",
        "up_category_info_test.go",
        "up_income_test.go",
        "up_info_video_test.go",
        "withdraw_test.go",
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
        "archive_income.go",
        "av_breach.go",
        "blacklist.go",
        "charge.go",
        "const.go",
        "format.go",
        "lottery.go",
        "service.go",
        "up_account.go",
        "up_category_info.go",
        "up_income.go",
        "up_info_video.go",
        "withdraw.go",
    ],
    importpath = "go-common/app/admin/main/growup/service/income",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/growup/conf",
        "//app/admin/main/growup/dao",
        "//app/admin/main/growup/dao/income",
        "//app/admin/main/growup/dao/message",
        "//app/admin/main/growup/dao/resource",
        "//app/admin/main/growup/model",
        "//app/admin/main/growup/model/income",
        "//app/admin/main/growup/service",
        "//library/database/sql",
        "//library/log",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
