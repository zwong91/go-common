package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "anchorReward",
    srcs = [
        "dao.cache.go",
        "dao.go",
        "mc.cache.go",
        "mc.go",
        "mysql.go",
        "redis.go",
    ],
    importpath = "go-common/app/service/live/xrewardcenter/dao/anchorReward",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/xrewardcenter/conf",
        "//app/service/live/xrewardcenter/model/anchorTask",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/orm",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
        "//library/time",
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

go_test(
    name = "anchorReward_test",
    srcs = [
        "dao.cache_test.go",
        "dao_test.go",
        "mc.cache_test.go",
        "mc_test.go",
        "mysql_test.go",
        "redis_test.go",
    ],
    embed = [":anchorReward"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/xrewardcenter/conf",
        "//app/service/live/xrewardcenter/model/anchorTask",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
