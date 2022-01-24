package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "timemachine",
    srcs = [
        "cache.go",
        "dao.cache.go",
        "dao.go",
        "hbase.go",
        "hbase_scan.go",
        "mc.cache.go",
    ],
    importpath = "go-common/app/interface/main/activity/dao/timemachine",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/model/timemachine",
        "//library/cache/memcache",
        "//library/database/hbase.v2:hbase_v2",
        "//library/log",
        "//library/stat/prom",
        "//library/sync/pipeline/fanout",
        "//vendor/github.com/tsuna/gohbase/hrpc",
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
    name = "timemachine_test",
    srcs = [
        "cache_test.go",
        "dao.cache_test.go",
        "dao_test.go",
        "hbase_test.go",
        "mc.cache_test.go",
    ],
    embed = [":timemachine"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/model/timemachine",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/github.com/tsuna/gohbase/hrpc",
    ],
)
