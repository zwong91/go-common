package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "associate_test.go",
        "business_test.go",
        "coupon_old_test.go",
        "coupon_test.go",
        "dao_test.go",
        "http_test.go",
        "memcache_test.go",
        "mysql_test.go",
        "oldvip_test.go",
        "order_test.go",
        "push_test.go",
        "redis_test.go",
        "resource_test.go",
        "sync_test.go",
        "vip_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/vip/conf",
        "//app/job/main/vip/model",
        "//library/database/sql",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "associate.go",
        "business.go",
        "coupon.go",
        "coupon_old.go",
        "dao.go",
        "http.go",
        "memcache.go",
        "mysql.go",
        "oldvip.go",
        "order.go",
        "push.go",
        "redis.go",
        "resource.go",
        "sync.go",
        "vip.go",
    ],
    importpath = "go-common/app/job/main/vip/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/vip/conf",
        "//app/job/main/vip/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/google/uuid",
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
