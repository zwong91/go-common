load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "ugc_test",
    srcs = [
        "dao_test.go",
        "databus_test.go",
        "del_arc_test.go",
        "del_upper_test.go",
        "del_video_test.go",
        "delete_test.go",
        "import_test.go",
        "manual_test.go",
        "media_cache_test.go",
        "passed_test.go",
        "search_con_test.go",
        "sync_video_test.go",
    ],
    embed = [":ugc"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/job/main/tv/model/ugc",
        "//app/service/main/archive/api",
        "//library/database/sql",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "ugc",
    srcs = [
        "dao.go",
        "databus.go",
        "del_arc.go",
        "del_upper.go",
        "del_video.go",
        "delete.go",
        "import.go",
        "manual.go",
        "media_cache.go",
        "passed.go",
        "report_cid.go",
        "search_con.go",
        "sync_arc.go",
        "sync_video.go",
    ],
    importpath = "go-common/app/job/main/tv/dao/ugc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/job/main/tv/dao/app",
        "//app/job/main/tv/model/common",
        "//app/job/main/tv/model/pgc",
        "//app/job/main/tv/model/ugc",
        "//app/service/main/archive/api",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
