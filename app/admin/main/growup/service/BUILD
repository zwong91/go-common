package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "authority_test.go",
        "blacklist_test.go",
        "cheat_test.go",
        "notice_test.go",
        "offlineactivity_test.go",
        "service_test.go",
        "tag_test.go",
        "trade_test.go",
        "up_info_test.go",
        "upload_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/growup/conf",
        "//app/admin/main/growup/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "activity.go",
        "authority.go",
        "auto.go",
        "banner.go",
        "blacklist.go",
        "budgetstatistics.go",
        "business.go",
        "cheat.go",
        "format.go",
        "notice.go",
        "offlineactivity.go",
        "service.go",
        "speical_award.go",
        "tag.go",
        "trade.go",
        "up_info.go",
        "upload.go",
        "white.go",
    ],
    importpath = "go-common/app/admin/main/growup/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/growup/conf",
        "//app/admin/main/growup/dao",
        "//app/admin/main/growup/dao/message",
        "//app/admin/main/growup/dao/resource",
        "//app/admin/main/growup/dao/shell",
        "//app/admin/main/growup/model",
        "//app/admin/main/growup/model/offlineactivity",
        "//app/admin/main/growup/util",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/jinzhu/gorm",
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
    srcs = [
        ":package-srcs",
        "//app/admin/main/growup/service/income:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
