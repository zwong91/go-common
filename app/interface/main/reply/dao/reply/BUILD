package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "reply_test",
    srcs = [
        "admin_test.go",
        "block_status_test.go",
        "business_test.go",
        "captcha_test.go",
        "config_test.go",
        "content_test.go",
        "credit_user_test.go",
        "cursor_test.go",
        "dao_test.go",
        "databus_test.go",
        "emoji_test.go",
        "memcache_test.go",
        "notice_test.go",
        "redis_test.go",
        "reply_test.go",
        "report_test.go",
        "subject_test.go",
    ],
    embed = [":reply"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/reply/conf",
        "//app/interface/main/reply/model/reply",
        "//library/database/sql",
        "//library/net/http/blademaster",
        "//library/queue/databus",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "reply",
    srcs = [
        "admin.go",
        "block_status.go",
        "business.go",
        "captcha.go",
        "config.go",
        "content.go",
        "credit_user.go",
        "cursor.go",
        "dao.go",
        "databus.go",
        "emoji.go",
        "memcache.go",
        "notice.go",
        "redis.go",
        "reply.go",
        "report.go",
        "subject.go",
    ],
    importpath = "go-common/app/interface/main/reply/dao/reply",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/reply/conf",
        "//app/interface/main/reply/model/reply",
        "//app/interface/main/reply/model/xreply",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/queue/databus",
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
