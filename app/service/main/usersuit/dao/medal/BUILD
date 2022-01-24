package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "medal_test",
    srcs = [
        "dao_test.go",
        "depend_test.go",
        "localcache_test.go",
        "memcache_test.go",
        "mysql_test.go",
    ],
    embed = [":medal"],
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
    name = "medal",
    srcs = [
        "dao.go",
        "depend.go",
        "localcache.go",
        "memcache.go",
        "mysql.go",
    ],
    importpath = "go-common/app/service/main/usersuit/dao/medal",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/usersuit/conf",
        "//app/service/main/usersuit/model",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
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
