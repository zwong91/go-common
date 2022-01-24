package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "archive_test",
    srcs = [
        "api_test.go",
        "biz_test.go",
        "dao_test.go",
        "fav_test.go",
        "history_test.go",
        "memcache_test.go",
        "music_test.go",
        "netsafe_test.go",
        "porder_test.go",
        "rpc_test.go",
        "staff_test.go",
        "viewpoint_test.go",
    ],
    embed = [":archive"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/archive",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "archive",
    srcs = [
        "api.go",
        "biz.go",
        "dao.cache.go",
        "dao.go",
        "fav.go",
        "history.go",
        "mc.cache.go",
        "memcache.go",
        "music.go",
        "netsafe.go",
        "porder.go",
        "rpc.go",
        "staff.go",
        "viewpoint.go",
    ],
    importpath = "go-common/app/interface/main/creative/dao/archive",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/tool",
        "//app/interface/main/creative/model/archive",
        "//app/interface/main/creative/model/music",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/stat/prom",
        "//library/sync/pipeline/fanout",
        "//library/xstr",
        "//vendor/golang.org/x/sync/singleflight",
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
