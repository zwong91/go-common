package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "activity_test.go",
        "authority_test.go",
        "banner_test.go",
        "blacklist_test.go",
        "block_test.go",
        "cheat_test.go",
        "credit_test.go",
        "dao_test.go",
        "expense_info_test.go",
        "notice_test.go",
        "offlineactivity_test.go",
        "special_award_test.go",
        "tag_test.go",
        "trade_test.go",
        "up_tag_income_test.go",
        "up_test.go",
        "upload_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/growup/conf",
        "//app/admin/main/growup/dao/shell",
        "//app/admin/main/growup/model",
        "//app/admin/main/growup/model/offlineactivity",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "activity.go",
        "authority.go",
        "banner.go",
        "blacklist.go",
        "block.go",
        "cheat.go",
        "credit.go",
        "dao.go",
        "expense_info.go",
        "notice.go",
        "offlineactivity.go",
        "special_award.go",
        "tag.go",
        "trade.go",
        "up.go",
        "up_tag_income.go",
        "upload.go",
    ],
    importpath = "go-common/app/admin/main/growup/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/growup/conf",
        "//app/admin/main/growup/dao/shell",
        "//app/admin/main/growup/model",
        "//app/admin/main/growup/model/offlineactivity",
        "//app/admin/main/up/util",
        "//library/database/orm",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/trace",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/jinzhu/gorm",
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
        "//app/admin/main/growup/dao/income:all-srcs",
        "//app/admin/main/growup/dao/message:all-srcs",
        "//app/admin/main/growup/dao/resource:all-srcs",
        "//app/admin/main/growup/dao/shell:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
