package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "up",
    srcs = [
        "dao.go",
        "mcn_account.go",
        "mcn_manage.go",
        "mcn_recommend.go",
        "memcache.go",
        "statistics.go",
        "up_base_Info.go",
    ],
    importpath = "go-common/app/admin/main/mcn/dao/up",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/mcn/conf",
        "//app/admin/main/mcn/model",
        "//app/interface/main/mcn/model/datamodel",
        "//app/interface/main/mcn/model/mcnmodel",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/time",
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
    name = "up_test",
    srcs = [
        "dao_test.go",
        "mcn_account_test.go",
        "mcn_manage_test.go",
        "mcn_recommend_test.go",
        "memcache_test.go",
        "statistics_test.go",
        "up_base_Info_test.go",
    ],
    embed = [":up"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/mcn/conf",
        "//app/admin/main/mcn/model",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)
