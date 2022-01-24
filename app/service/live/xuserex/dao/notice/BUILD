package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "notice",
    srcs = [
        "dao.cache.go",
        "mc.cache.go",
        "roomNotice.go",
    ],
    importpath = "go-common/app/service/live/xuserex/dao/notice",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/resource/sdk",
        "//app/service/live/xuser/api/grpc/v1:grpc",
        "//app/service/live/xuserex/api/grpc/v1:grpc",
        "//app/service/live/xuserex/conf",
        "//app/service/live/xuserex/model/roomNotice",
        "//library/cache/memcache",
        "//library/database/hbase.v2:hbase_v2",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
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
    name = "notice_test",
    srcs = [
        "dao.cache_test.go",
        "mc.cache_test.go",
        "roomNotice_test.go",
    ],
    embed = [":notice"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/xuserex/conf",
        "//app/service/live/xuserex/model/roomNotice",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
