package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "bfs_test.go",
        "dao_test.go",
        "mask_test.go",
        "mc_special_test.go",
        "mc_subtitle_test.go",
        "memcache_seg_test.go",
        "memcache_test.go",
        "mysql_dm_special_test.go",
        "mysql_test.go",
        "rank_list_test.go",
        "redis_mask_test.go",
        "redis_rct_test.go",
        "redis_seg_test.go",
        "redis_task_test.go",
        "redis_test.go",
        "task_test.go",
        "transfer_test.go",
        "videoup_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/dm2/conf",
        "//app/job/main/dm2/model",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "bfs.go",
        "bnj.go",
        "dao.go",
        "mask.go",
        "mc_special.go",
        "mc_subtitle.go",
        "memcache.go",
        "memcache_seg.go",
        "mysql.go",
        "mysql_dm_special.go",
        "rank_list.go",
        "redis.go",
        "redis_mask.go",
        "redis_rct.go",
        "redis_seg.go",
        "redis_task.go",
        "task.go",
        "transfer.go",
        "videoup.go",
    ],
    importpath = "go-common/app/job/main/dm2/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/dm2/conf",
        "//app/job/main/dm2/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/bfs",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/sync/errgroup",
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
