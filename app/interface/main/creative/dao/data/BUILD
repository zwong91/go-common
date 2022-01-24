package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "data",
    srcs = [
        "api.go",
        "archive.go",
        "archive_play.go",
        "article.go",
        "dao.go",
        "fan.go",
        "memcache.go",
        "overview.go",
        "overview_mc.go",
        "table.go",
    ],
    importpath = "go-common/app/interface/main/creative/dao/data",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up/util/hbaseutil",
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/data",
        "//app/interface/openplatform/article/model",
        "//library/cache/memcache",
        "//library/database/hbase.v2:hbase_v2",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/github.com/tsuna/gohbase/hrpc",
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

go_test(
    name = "data_test",
    srcs = [
        "api_test.go",
        "archive_play_test.go",
        "archive_test.go",
        "article_test.go",
        "dao_test.go",
        "fan_test.go",
        "memcache_test.go",
        "overview_mc_test.go",
        "overview_test.go",
    ],
    embed = [":data"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/data",
        "//app/interface/openplatform/article/model",
        "//library/cache/memcache",
        "//library/database/hbase.v2:hbase_v2",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/github.com/tsuna/gohbase/hrpc",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)
