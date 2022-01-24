package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "archive_test",
    srcs = [
        "archive_test.go",
        "dao_test.go",
        "delay_test.go",
        "desc_format_test.go",
        "flow_test.go",
        "history_test.go",
        "netsafe_test.go",
        "new_video_test.go",
        "porder_test.go",
        "redis_test.go",
        "report_test.go",
        "staff_test.go",
        "type_test.go",
        "video_test.go",
    ],
    embed = [":archive"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/videoup/conf",
        "//app/service/main/videoup/model/archive",
        "//library/database/sql",
        "//library/time",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/davecgh/go-spew/spew",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "archive",
    srcs = [
        "archive.go",
        "dao.cache.go",
        "dao.go",
        "delay.go",
        "desc_format.go",
        "flow.go",
        "forbid.go",
        "history.go",
        "mc.cache.go",
        "netsafe.go",
        "new_video.go",
        "porder.go",
        "redis.go",
        "relation.go",
        "report.go",
        "staff.go",
        "staff_apply.go",
        "type.go",
        "video.go",
    ],
    importpath = "go-common/app/service/main/videoup/dao/archive",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/videoup/conf",
        "//app/service/main/videoup/model/archive",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/net/metadata",
        "//library/stat/prom",
        "//library/sync/pipeline/fanout",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/dgryski/go-farm",
        "//vendor/github.com/pkg/errors",
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
