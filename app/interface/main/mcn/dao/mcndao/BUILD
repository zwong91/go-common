package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "mcndao",
    srcs = [
        "crm.go",
        "dao.cache.go",
        "dao.go",
        "data.go",
        "mc.cache.go",
        "memcahe.go",
        "rank.go",
        "recommend_pool.go",
        "sign.go",
    ],
    importpath = "go-common/app/interface/main/mcn/dao/mcndao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/mcn/model",
        "//app/admin/main/up/util",
        "//app/admin/main/up/util/mathutil",
        "//app/interface/main/mcn/conf",
        "//app/interface/main/mcn/dao/cache",
        "//app/interface/main/mcn/dao/global",
        "//app/interface/main/mcn/model",
        "//app/interface/main/mcn/model/mcnmodel",
        "//app/interface/main/mcn/tool/validate",
        "//app/service/main/archive/api",
        "//library/cache/memcache",
        "//library/ecode",
        "//library/log",
        "//library/stat/prom",
        "//library/sync/pipeline/fanout",
        "//vendor/github.com/bluele/gcache",
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
    name = "mcndao_test",
    srcs = [
        "crm_test.go",
        "dao.cache_test.go",
        "dao_test.go",
        "data_test.go",
        "mc.cache_test.go",
        "memcahe_test.go",
        "rank_test.go",
        "recommend_pool_test.go",
        "sign_test.go",
    ],
    embed = [":mcndao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/mcn/conf",
        "//app/interface/main/mcn/dao/global",
        "//app/interface/main/mcn/model",
        "//app/interface/main/mcn/model/mcnmodel",
        "//library/ecode",
        "//vendor/github.com/bluele/gcache",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
