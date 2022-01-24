package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "business",
    srcs = [
        "aegis_resource.go",
        "archive_video_relation.go",
        "avr_archive.go",
        "dm_date.go",
        "log.go",
    ],
    importpath = "go-common/app/job/main/search/dao/business",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/search/conf",
        "//app/job/main/search/dao",
        "//app/job/main/search/model",
        "//library/database/sql",
        "//library/log",
        "//library/log/infoc",
        "//library/queue/databus",
        "//library/xstr",
        "//vendor/gopkg.in/olivere/elastic.v5:elastic_v5",
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
    name = "business_test",
    srcs = ["business_test.go"],
    embed = [":business"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/search/conf",
        "//app/job/main/search/dao",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
