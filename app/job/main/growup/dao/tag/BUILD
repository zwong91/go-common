package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "tag_test",
    srcs = [
        "activity_test.go",
        "archive_charge_test.go",
        "charge_ratio_test.go",
        "dao_test.go",
        "extra_test.go",
        "income_test.go",
        "tag_info_test.go",
        "tag_up_info_test.go",
    ],
    embed = [":tag"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/growup/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "tag",
    srcs = [
        "activity.go",
        "archive_charge.go",
        "charge_ratio.go",
        "dao.go",
        "extra.go",
        "income.go",
        "tag_info.go",
        "tag_up_info.go",
        "up_tag_income.go",
    ],
    importpath = "go-common/app/job/main/growup/dao/tag",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/growup/conf",
        "//app/job/main/growup/model",
        "//app/job/main/growup/model/tag",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
