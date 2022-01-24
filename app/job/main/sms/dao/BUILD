package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = ["user_test.go"],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/sms/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "user.go",
        "wechat.go",
    ],
    importpath = "go-common/app/job/main/sms/dao",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/sms/conf",
        "//app/job/main/sms/model",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
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
        "//app/job/main/sms/dao/chuanglan:all-srcs",
        "//app/job/main/sms/dao/mengwang:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
