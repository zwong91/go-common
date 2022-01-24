package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "like_test",
    srcs = [
        "action_test.go",
        "dao_test.go",
        "extend_test.go",
        "like_test.go",
        "match_test.go",
        "redis_test.go",
        "subject_test.go",
        "web_data_test.go",
    ],
    embed = [":like"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/activity/conf",
        "//app/job/main/activity/model/like",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "like",
    srcs = [
        "action.go",
        "dao.go",
        "extend.go",
        "like.go",
        "match.go",
        "redis.go",
        "subject.go",
        "web_data.go",
    ],
    importpath = "go-common/app/job/main/activity/dao/like",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/activity/model",
        "//app/job/main/activity/conf",
        "//app/job/main/activity/model/like",
        "//app/job/main/activity/model/match",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/elastic",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/xstr",
        "//vendor/github.com/pkg/errors",
        "//vendor/golang.org/x/net/context",
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
