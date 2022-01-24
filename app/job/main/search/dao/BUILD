package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "app_multiple_databus_test.go",
        "app_single_test.go",
        "config_offset_test.go",
        "dao_test.go",
        "es_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/search/conf",
        "//app/job/main/search/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "app_databus.go",
        "app_multiple.go",
        "app_multiple_databus.go",
        "app_single.go",
        "config_asset.go",
        "config_attr.go",
        "config_business.go",
        "config_offset.go",
        "dao.go",
        "es.go",
        "hbase.go",
        "sms.go",
    ],
    importpath = "go-common/app/job/main/search/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/search/conf",
        "//app/job/main/search/model",
        "//library/database/sql",
        "//library/log",
        "//library/log/infoc",
        "//library/queue/databus",
        "//library/stat/prom",
        "//library/time",
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
    srcs = [
        ":package-srcs",
        "//app/job/main/search/dao/base:all-srcs",
        "//app/job/main/search/dao/business:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
