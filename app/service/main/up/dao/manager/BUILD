package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "manager_test",
    srcs = [
        "dao.cache_test.go",
        "dao_test.go",
        "mc.cache_test.go",
        "up_group_test.go",
        "up_mc_test.go",
        "up_test.go",
    ],
    embed = [":manager"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/up/api/v1:api",
        "//app/service/main/up/conf",
        "//app/service/main/up/dao",
        "//app/service/main/up/dao/global",
        "//app/service/main/up/model",
        "//library/net/http/blademaster",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "manager",
    srcs = [
        "dao.cache.go",
        "dao.go",
        "mc.cache.go",
        "up.go",
        "up_group.go",
        "up_mc.go",
    ],
    importpath = "go-common/app/service/main/up/dao/manager",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/up/api/v1:api",
        "//app/service/main/up/conf",
        "//app/service/main/up/dao",
        "//app/service/main/up/dao/global",
        "//app/service/main/up/model",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
        "//library/xstr",
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
