package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "blacklist_test.go",
        "cheat_test.go",
        "combinemails_test.go",
        "data_test.go",
        "httpClient_test.go",
        "income_test.go",
        "recompute_test.go",
        "service_test.go",
        "tag_test.go",
        "tagincome_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/growup/conf",
        "//app/job/main/growup/service/ctrl",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "auto.go",
        "blacklist.go",
        "cheat.go",
        "combinemails.go",
        "creative_activity.go",
        "creative_budget.go",
        "data.go",
        "databus.go",
        "format.go",
        "httpClient.go",
        "income.go",
        "lottery.go",
        "recompute.go",
        "sendincome.go",
        "service.go",
        "signedups.go",
        "tag.go",
        "tagincome.go",
        "task_status.go",
        "tax.go",
        "tool.go",
        "topten.go",
        "up.go",
        "up_bill.go",
        "upload.go",
        "video.go",
    ],
    importpath = "go-common/app/job/main/growup/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/growup/conf",
        "//app/job/main/growup/dao",
        "//app/job/main/growup/dao/charge",
        "//app/job/main/growup/dao/dataplatform",
        "//app/job/main/growup/dao/email",
        "//app/job/main/growup/dao/income",
        "//app/job/main/growup/dao/tag",
        "//app/job/main/growup/model",
        "//app/job/main/growup/model/income",
        "//app/job/main/growup/service/ctrl",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
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
        "//app/job/main/growup/service/charge:all-srcs",
        "//app/job/main/growup/service/ctrl:all-srcs",
        "//app/job/main/growup/service/income:all-srcs",
        "//app/job/main/growup/service/tag:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
