package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.cache.go",
        "dao.go",
        "grpc.go",
        "http.go",
        "mc.cache.go",
        "mysql.go",
        "redis.go",
    ],
    importpath = "go-common/app/service/main/ugcpay/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/archive/api",
        "//app/service/main/ugcpay/conf",
        "//app/service/main/ugcpay/model",
        "//library/cache",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/stat/prom",
        "//vendor/github.com/pkg/errors",
        "//vendor/golang.org/x/sync/singleflight",
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
    name = "dao_test",
    srcs = [
        "dao.cache_test.go",
        "dao_test.go",
        "grpc_test.go",
        "http_test.go",
        "mc.cache_test.go",
        "mysql_test.go",
        "redis_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/ugcpay/conf",
        "//app/service/main/ugcpay/model",
        "//library/database/sql",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
