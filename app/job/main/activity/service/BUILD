package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "match_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/activity/conf",
        "//app/job/main/activity/model/like",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "action.go",
        "archive.go",
        "bnj.go",
        "kfc.go",
        "like.go",
        "match.go",
        "service.go",
        "subject.go",
    ],
    importpath = "go-common/app/job/main/activity/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/activity/model",
        "//app/interface/main/activity/model/like",
        "//app/interface/main/activity/rpc/client",
        "//app/interface/openplatform/article/model",
        "//app/interface/openplatform/article/rpc/client",
        "//app/job/main/activity/conf",
        "//app/job/main/activity/dao/bnj",
        "//app/job/main/activity/dao/dm",
        "//app/job/main/activity/dao/kfc",
        "//app/job/main/activity/dao/like",
        "//app/job/main/activity/model/bnj",
        "//app/job/main/activity/model/kfc",
        "//app/job/main/activity/model/like",
        "//app/job/main/activity/model/match",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/coin/api/gorpc",
        "//app/service/main/coin/model",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//vendor/github.com/robfig/cron",
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
