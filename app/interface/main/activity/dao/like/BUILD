package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "like_test",
    srcs = [
        "achievements_test.go",
        "act_mission_test.go",
        "act_test.go",
        "action_test.go",
        "cache_test.go",
        "content_test.go",
        "dao.cache_test.go",
        "dao_test.go",
        "extend_test.go",
        "like_test.go",
        "lottery_test.go",
        "match_redis_test.go",
        "match_test.go",
        "mc.cache_test.go",
        "memcached_test.go",
        "mission_group_test.go",
        "protocol_test.go",
        "redis_test.go",
        "subject_test.go",
        "user_achievements_test.go",
    ],
    embed = [":like"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/model/like",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "like",
    srcs = [
        "achievements.go",
        "act.go",
        "act_mission.go",
        "action.go",
        "cache.go",
        "content.go",
        "dao.cache.go",
        "dao.go",
        "extend.go",
        "like.go",
        "lottery.go",
        "match.go",
        "match_redis.go",
        "mc.cache.go",
        "memcached.go",
        "mission_group.go",
        "protocol.go",
        "redis.go",
        "subject.go",
        "user_achievements.go",
    ],
    importpath = "go-common/app/interface/main/activity/dao/like",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/model/like",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/elastic",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/stat/prom",
        "//library/sync/pipeline/fanout",
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
