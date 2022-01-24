package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "tag_test",
    srcs = [
        "service_test.go",
        "tag_test.go",
        "tag_up_info_test.go",
        "up_tag_income_test.go",
    ],
    embed = [":tag"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/growup/conf",
        "//app/job/main/growup/service/ctrl",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "tag",
    srcs = [
        "activity.go",
        "archive_charge.go",
        "archive_tag.go",
        "category.go",
        "charge_ratio.go",
        "extra.go",
        "income.go",
        "service.go",
        "tag.go",
        "tag_up_info.go",
        "up_tag.go",
        "up_tag_income.go",
    ],
    importpath = "go-common/app/job/main/growup/service/tag",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/growup/conf",
        "//app/job/main/growup/dao/email",
        "//app/job/main/growup/dao/tag",
        "//app/job/main/growup/model/tag",
        "//app/job/main/growup/service",
        "//app/job/main/growup/service/ctrl",
        "//library/database/sql",
        "//library/log",
        "//library/time",
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
