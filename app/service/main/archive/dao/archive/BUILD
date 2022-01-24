package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "archive_test",
    srcs = [
        "addit_test.go",
        "archive_mc_pb_test.go",
        "archive_test.go",
        "dao_test.go",
        "player_test.go",
        "region_rds_test.go",
        "region_test.go",
        "report_result_test.go",
        "stat_mc_pb_test.go",
        "stat_pb_test.go",
        "stat_test.go",
        "type_test.go",
        "upper_rds_test.go",
        "upper_test.go",
        "video_mc_pb_test.go",
        "video_test.go",
    ],
    embed = [":archive"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/api",
        "//app/service/main/archive/conf",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "archive",
    srcs = [
        "addit.go",
        "archive.go",
        "archive_mc_pb.go",
        "dao.go",
        "player.go",
        "region.go",
        "region_rds.go",
        "report_result.go",
        "stat.go",
        "stat_mc_pb.go",
        "stat_pb.go",
        "type.go",
        "upper.go",
        "upper_rds.go",
        "video.go",
        "video_mc_pb.go",
    ],
    importpath = "go-common/app/service/main/archive/dao/archive",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/archive/model/databus",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/archive/conf",
        "//app/service/main/archive/model/archive",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/queue/databus",
        "//library/stat/prom",
        "//library/sync/errgroup",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/pkg/errors",
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
