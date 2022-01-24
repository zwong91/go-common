package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "archive_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/videoup/conf",
        "//app/interface/main/videoup/model/archive",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "allow.go",
        "app.go",
        "check.go",
        "check_staff.go",
        "client.go",
        "creator.go",
        "deal.go",
        "filter.go",
        "geetest.go",
        "infoc.go",
        "log.go",
        "pay.go",
        "pre.go",
        "service.go",
        "subtitle.go",
        "web.go",
    ],
    importpath = "go-common/app/interface/main/videoup/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/model/tag",
        "//app/interface/main/videoup/conf",
        "//app/interface/main/videoup/dao/account",
        "//app/interface/main/videoup/dao/archive",
        "//app/interface/main/videoup/dao/bfs",
        "//app/interface/main/videoup/dao/creative",
        "//app/interface/main/videoup/dao/dynamic",
        "//app/interface/main/videoup/dao/elec",
        "//app/interface/main/videoup/dao/filter",
        "//app/interface/main/videoup/dao/geetest",
        "//app/interface/main/videoup/dao/mission",
        "//app/interface/main/videoup/dao/order",
        "//app/interface/main/videoup/dao/pay",
        "//app/interface/main/videoup/dao/subtitle",
        "//app/interface/main/videoup/dao/tag",
        "//app/interface/main/videoup/model/archive",
        "//app/interface/main/videoup/model/geetest",
        "//app/interface/main/videoup/model/mission",
        "//app/interface/main/videoup/model/porder",
        "//app/service/main/account/api",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//library/net/metadata",
        "//library/queue/databus/report",
        "//library/sync/errgroup",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/microcosm-cc/bluemonday",
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
