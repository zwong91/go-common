package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "danmu",
    srcs = [
        "adv_danmu.go",
        "assist.go",
        "dao.go",
        "dm.go",
        "protect.go",
        "report.go",
    ],
    importpath = "go-common/app/interface/main/creative/dao/danmu",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/tool",
        "//app/interface/main/creative/model/danmu",
        "//app/interface/main/dm/model",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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

go_test(
    name = "danmu_test",
    srcs = [
        "adv_danmu_test.go",
        "assist_test.go",
        "dao_test.go",
        "dm_test.go",
        "protect_test.go",
        "report_test.go",
    ],
    embed = [":danmu"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/danmu",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)
