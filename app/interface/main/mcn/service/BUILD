package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "account.go",
        "api.go",
        "cmd.go",
        "creative.go",
        "data.go",
        "mcn.go",
        "msg.go",
        "rank.go",
        "service.go",
        "uniquecheck.go",
        "upload.go",
    ],
    importpath = "go-common/app/interface/main/mcn/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/mcn/model",
        "//app/interface/main/mcn/conf",
        "//app/interface/main/mcn/dao/bfs",
        "//app/interface/main/mcn/dao/cache",
        "//app/interface/main/mcn/dao/datadao",
        "//app/interface/main/mcn/dao/global",
        "//app/interface/main/mcn/dao/mcndao",
        "//app/interface/main/mcn/dao/msg",
        "//app/interface/main/mcn/model",
        "//app/interface/main/mcn/model/datamodel",
        "//app/interface/main/mcn/model/mcnmodel",
        "//app/interface/main/mcn/tool/worker",
        "//app/service/main/account/api",
        "//app/service/main/member/api",
        "//app/service/main/member/model",
        "//app/service/main/member/model/block",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/time",
        "//vendor/github.com/bluele/gcache",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "service_test",
    srcs = [
        "account_test.go",
        "cmd_test.go",
        "mcn_test.go",
        "msg_test.go",
        "rank_test.go",
        "service_test.go",
        "uniquecheck_test.go",
        "upload_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/mcn/model",
        "//app/interface/main/mcn/conf",
        "//app/interface/main/mcn/model",
        "//app/interface/main/mcn/model/mcnmodel",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
