package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "allowance_test.go",
        "captcha_api_test.go",
        "code_test.go",
        "dao_test.go",
        "http_test.go",
        "memcache_test.go",
        "mysql_test.go",
        "prize_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/coupon/conf",
        "//app/service/main/coupon/model",
        "//library/database/sql",
        "//library/ecode",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "allowance.go",
        "captcha_api.go",
        "code.go",
        "dao.go",
        "http.go",
        "memcache.go",
        "mysql.go",
        "prize.go",
    ],
    importpath = "go-common/app/service/main/coupon/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/coupon/conf",
        "//app/service/main/coupon/model",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/stat/prom",
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
