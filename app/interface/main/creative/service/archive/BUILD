package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "archive_test",
    srcs = ["service_test.go"],
    embed = [":archive"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/archive",
        "//app/interface/main/creative/model/order",
        "//app/interface/main/creative/service",
        "//vendor/github.com/davecgh/go-spew/spew",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "archive",
    srcs = [
        "archive.go",
        "biz.go",
        "fav.go",
        "history.go",
        "mission.go",
        "myinfo.go",
        "netsafe.go",
        "order.go",
        "porder.go",
        "rules.go",
        "service.go",
        "staff.go",
        "tag.go",
        "viewpoint.go",
    ],
    importpath = "go-common/app/interface/main/creative/service/archive",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/videoup/model/archive",
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/account",
        "//app/interface/main/creative/dao/activity",
        "//app/interface/main/creative/dao/appeal",
        "//app/interface/main/creative/dao/archive",
        "//app/interface/main/creative/dao/coin",
        "//app/interface/main/creative/dao/game",
        "//app/interface/main/creative/dao/order",
        "//app/interface/main/creative/dao/search",
        "//app/interface/main/creative/dao/tag",
        "//app/interface/main/creative/dao/template",
        "//app/interface/main/creative/model/activity",
        "//app/interface/main/creative/model/appeal",
        "//app/interface/main/creative/model/archive",
        "//app/interface/main/creative/model/game",
        "//app/interface/main/creative/model/order",
        "//app/interface/main/creative/model/search",
        "//app/interface/main/creative/model/tag",
        "//app/interface/main/creative/service",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//library/ecode",
        "//library/log",
        "//library/stat/prom",
        "//library/sync/errgroup",
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
