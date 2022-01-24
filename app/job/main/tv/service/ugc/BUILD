load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "ugc_test",
    srcs = [
        "import_test.go",
        "sea_con_test.go",
        "sync_video_test.go",
        "ugc_test.go",
    ],
    embed = [":ugc"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/job/main/tv/model/ugc",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "ugc",
    srcs = [
        "databus.go",
        "del_arc.go",
        "del_video.go",
        "full_refresh.go",
        "import.go",
        "manual.go",
        "media_cache.go",
        "pick.go",
        "report_cid.go",
        "sea_con.go",
        "sync_arc.go",
        "sync_video.go",
        "tool.go",
        "ugc.go",
        "upper.go",
        "view.go",
        "zone_index.go",
    ],
    importpath = "go-common/app/job/main/tv/service/ugc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/job/main/tv/dao/app",
        "//app/job/main/tv/dao/archive",
        "//app/job/main/tv/dao/cms",
        "//app/job/main/tv/dao/ftp",
        "//app/job/main/tv/dao/lic",
        "//app/job/main/tv/dao/playurl",
        "//app/job/main/tv/dao/ugc",
        "//app/job/main/tv/dao/upper",
        "//app/job/main/tv/model/pgc",
        "//app/job/main/tv/model/ugc",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//library/time",
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
