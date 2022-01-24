package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "business_test.go",
        "channel_category_test.go",
        "channel_group_test.go",
        "channel_mc_test.go",
        "channel_rule_test.go",
        "channel_test.go",
        "dao_test.go",
        "es_test.go",
        "grpc_test.go",
        "http_test.go",
        "limit_test.go",
        "rank_test.go",
        "region_tag_archive_redis_test.go",
        "relation_mc_test.go",
        "relation_redis_test.go",
        "relation_test.go",
        "report_test.go",
        "resource_log_test.go",
        "synonym_test.go",
        "tag_mc_test.go",
        "tag_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/tag/conf",
        "//app/admin/main/tag/model",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "business.go",
        "channel.go",
        "channel_category.go",
        "channel_group.go",
        "channel_mc.go",
        "channel_rule.go",
        "dao.go",
        "es.go",
        "grpc.go",
        "http.go",
        "limit.go",
        "rank.go",
        "region_tag_archive_redis.go",
        "relation.go",
        "relation_mc.go",
        "relation_redis.go",
        "report.go",
        "resource_log.go",
        "synonym.go",
        "tag.go",
        "tag_mc.go",
    ],
    importpath = "go-common/app/admin/main/tag/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/tag/conf",
        "//app/admin/main/tag/model",
        "//app/service/main/account/api",
        "//app/service/main/archive/model/archive",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/elastic",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
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
