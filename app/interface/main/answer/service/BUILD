package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "answer_test.go",
        "answercheck_test.go",
        "common_test.go",
        "geetest_test.go",
        "service_test.go",
        "user_info_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/answer/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "answer.go",
        "answercheck.go",
        "common.go",
        "geetest.go",
        "rank_share.go",
        "service.go",
        "user_info.go",
    ],
    importpath = "go-common/app/interface/main/answer/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/answer/conf",
        "//app/interface/main/answer/dao",
        "//app/interface/main/answer/dao/account",
        "//app/interface/main/answer/dao/geetest",
        "//app/interface/main/answer/model",
        "//app/service/main/account/api",
        "//app/service/main/member/api/gorpc",
        "//app/service/main/member/model",
        "//library/ecode",
        "//library/log",
        "//library/log/anticheat",
        "//library/log/infoc",
        "//library/net/metadata",
        "//library/queue/databus/report",
        "//library/stat/prom",
        "//library/sync/pipeline/fanout",
        "//library/text/translate/chinese",
        "//vendor/github.com/pkg/errors",
        "//vendor/golang.org/x/net/context",
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
