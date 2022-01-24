package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "pgc_test",
    srcs = [
        "sea_con_test.go",
        "service_test.go",
        "sync_mc_test.go",
        "zone_index_test.go",
    ],
    embed = [":pgc"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/tv/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "pgc",
    srcs = [
        "del_cont.go",
        "del_season.go",
        "filter.go",
        "license.go",
        "sea_con.go",
        "search_sug.go",
        "service.go",
        "sync_ep.go",
        "sync_mc.go",
        "sync_retry.go",
        "sync_season.go",
        "zone_index.go",
    ],
    importpath = "go-common/app/job/main/tv/service/pgc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/job/main/tv/dao/app",
        "//app/job/main/tv/dao/cms",
        "//app/job/main/tv/dao/ftp",
        "//app/job/main/tv/dao/lic",
        "//app/job/main/tv/dao/playurl",
        "//app/job/main/tv/model/common",
        "//app/job/main/tv/model/pgc",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//library/time",
        "//library/xstr",
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
