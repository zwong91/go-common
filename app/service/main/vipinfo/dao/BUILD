package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "cache.go",
        "dao.cache.go",
        "dao.go",
        "mc.go",
        "mysql.go",
    ],
    importpath = "go-common/app/service/main/vipinfo/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/vipinfo/conf",
        "//app/service/main/vipinfo/model",
        "//library/cache",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/net/metadata",
        "//library/stat/prom",
        "//library/sync/errgroup",
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

go_test(
    name = "dao_test",
    srcs = [
        "dao.cache_test.go",
        "dao_test.go",
        "mc_test.go",
        "mysql_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/vipinfo/conf",
        "//app/service/main/vipinfo/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
