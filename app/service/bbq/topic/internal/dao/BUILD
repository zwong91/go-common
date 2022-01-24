package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "common.go",
        "dao.cache.go",
        "dao.go",
        "extension.go",
        "topic.go",
        "topic_video.go",
    ],
    importpath = "go-common/app/service/bbq/topic/internal/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/topic/api",
        "//app/service/bbq/topic/internal/model",
        "//library/cache/redis",
        "//library/conf/paladin",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/stat/prom",
        "//library/sync/errgroup",
        "//library/sync/errgroup.v2:errgroup_v2",
        "//library/sync/pipeline/fanout",
        "//library/time",
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

go_test(
    name = "dao_test",
    srcs = [
        "common_test.go",
        "dao.cache_test.go",
        "dao_test.go",
        "extension_test.go",
        "topic_test.go",
        "topic_video_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/service/bbq/topic/api",
        "//app/service/bbq/topic/internal/model",
        "//library/conf/paladin",
        "//library/ecode",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
