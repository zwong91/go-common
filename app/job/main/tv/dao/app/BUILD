package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "app_test",
    srcs = [
        "cont_del_test.go",
        "cont_sync_test.go",
        "dao_test.go",
        "media_cache_test.go",
        "new_test.go",
        "playurl_test.go",
        "redis_test.go",
        "search_con_test.go",
        "search_sug_test.go",
        "season_del_test.go",
        "season_sync_test.go",
        "sync_mc_test.go",
        "sync_retry_test.go",
        "zone_index_test.go",
    ],
    embed = [":app"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/job/main/tv/model/common",
        "//app/job/main/tv/model/pgc",
        "//library/database/sql",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "app",
    srcs = [
        "cont_del.go",
        "cont_sync.go",
        "dao.go",
        "media_cache.go",
        "playurl.go",
        "redis.go",
        "search_con.go",
        "search_sug.go",
        "season_del.go",
        "season_sync.go",
        "sync_mc.go",
        "sync_retry.go",
        "zone_index.go",
    ],
    importpath = "go-common/app/job/main/tv/dao/app",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/job/main/tv/model/common",
        "//app/job/main/tv/model/pgc",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
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
