package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "grpc.go",
        "localcache.go",
        "mc.cache.go",
        "mc.extra.go",
        "mysql.go",
    ],
    importpath = "go-common/app/service/main/ugcpay-rank/internal/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/account/api",
        "//app/service/main/account/model",
        "//app/service/main/ugcpay-rank/internal/conf",
        "//app/service/main/ugcpay-rank/internal/model",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/log",
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

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "grpc_test.go",
        "localcache_test.go",
        "mc.cache_test.go",
        "mc.extra_test.go",
        "mysql_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/ugcpay-rank/internal/conf",
        "//app/service/main/ugcpay-rank/internal/model",
        "//library/cache/memcache",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
