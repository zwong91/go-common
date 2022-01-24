load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

package(default_visibility = ["//visibility:public"])

go_test(
    name = "dao_test",
    srcs = [
        "bfs_test.go",
        "dao_test.go",
        "face_history_test.go",
        "member_test.go",
        "merak_test.go",
        "message_test.go",
        "monitor_test.go",
        "official_test.go",
        "passport_test.go",
        "realname_test.go",
        "redis_test.go",
        "review_test.go",
        "search_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/member/conf",
        "//app/admin/main/member/model",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/github.com/tsuna/gohbase/hrpc",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "bfs.go",
        "dao.go",
        "face_history.go",
        "member.go",
        "merak.go",
        "message.go",
        "monitor.go",
        "official.go",
        "passport.go",
        "realname.go",
        "realname_old.go",
        "redis.go",
        "review.go",
        "search.go",
    ],
    importpath = "go-common/app/admin/main/member/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/member/conf",
        "//app/admin/main/member/dao/block",
        "//app/admin/main/member/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/elastic",
        "//library/database/hbase.v2:hbase_v2",
        "//library/database/orm",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/ip",
        "//library/net/metadata",
        "//library/queue/databus",
        "//library/sync/errgroup",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/pkg/errors",
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
    srcs = [
        ":package-srcs",
        "//app/admin/main/member/dao/block:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
