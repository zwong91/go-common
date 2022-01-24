package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "pendant_test",
    srcs = [
        "dao_test.go",
        "local_equip_test.go",
        "memcache_test.go",
        "mysql_test.go",
        "pay_test.go",
        "redis_test.go",
        "vip_test.go",
    ],
    embed = [":pendant"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/usersuit/conf",
        "//app/service/main/usersuit/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "pendant",
    srcs = [
        "dao.go",
        "local_equip.go",
        "memcache.go",
        "mysql.go",
        "pay.go",
        "redis.go",
        "vip.go",
    ],
    importpath = "go-common/app/service/main/usersuit/dao/pendant",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/usersuit/conf",
        "//app/service/main/usersuit/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
        "//library/xstr",
        "//vendor/github.com/bluele/gcache",
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
