package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "business_test.go",
        "channel_test.go",
        "limit_test.go",
        "rank_test.go",
        "relation_test.go",
        "report_test.go",
        "resource_log_test.go",
        "resource_test.go",
        "service_test.go",
        "synonym_test.go",
        "tag_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/tag/conf",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "business.go",
        "channel.go",
        "channel_category.go",
        "channel_rule.go",
        "limit.go",
        "rank.go",
        "relation.go",
        "report.go",
        "resource.go",
        "resource_log.go",
        "service.go",
        "synonym.go",
        "tag.go",
    ],
    importpath = "go-common/app/admin/main/tag/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/tag/conf",
        "//app/admin/main/tag/dao",
        "//app/admin/main/tag/model",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
