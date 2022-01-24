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
        "callback_test.go",
        "challenge_test.go",
        "dao_test.go",
        "event_test.go",
        "extra_test.go",
        "group_test.go",
        "memcache_test.go",
        "message_test.go",
        "redis_test.go",
        "search2_test.go",
        "tag_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/workflow/model",
        "//app/admin/main/workflow/model/param",
        "//app/admin/main/workflow/model/search",
        "//library/conf/paladin",
        "//library/ecode",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "business.go",
        "callback.go",
        "challenge.go",
        "dao.go",
        "event.go",
        "extra.go",
        "group.go",
        "memcache.go",
        "message.go",
        "redis.go",
        "search2.go",
        "tag.go",
    ],
    importpath = "go-common/app/admin/main/workflow/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/workflow/model",
        "//app/admin/main/workflow/model/manager",
        "//app/admin/main/workflow/model/param",
        "//app/admin/main/workflow/model/search",
        "//app/interface/main/credit/model",
        "//app/job/main/member/model/block",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/member/model/block",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/conf/paladin",
        "//library/database/elastic",
        "//library/database/orm",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/rpc/warden",
        "//library/sync/errgroup",
        "//library/xstr",
        "//vendor/github.com/jinzhu/gorm",
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
