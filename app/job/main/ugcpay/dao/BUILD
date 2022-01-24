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
        "elec_mc.go",
        "elec_mysql.go",
        "grpc.go",
        "http.go",
        "mc.go",
        "mysql.go",
        "redis.go",
    ],
    importpath = "go-common/app/job/main/ugcpay/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/ugcpay/conf",
        "//app/job/main/ugcpay/model",
        "//app/service/main/ugcpay-rank/api/v1:api",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/stat/prom",
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
        "elec_mc_test.go",
        "elec_mysql_test.go",
        "grpc_test.go",
        "http_test.go",
        "mc_test.go",
        "mysql_test.go",
        "redis_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/ugcpay/conf",
        "//app/job/main/ugcpay/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
