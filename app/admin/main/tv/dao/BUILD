load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "archive_test.go",
        "audit_result_test.go",
        "dao_test.go",
        "full_test.go",
        "modules_test.go",
        "panel_test.go",
        "pgc_cond_test.go",
        "playurl_test.go",
        "region_test.go",
        "sear_inter_test.go",
        "upbfs_test.go",
        "upper_test.go",
        "user_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/tv/conf",
        "//app/admin/main/tv/model",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "archive.go",
        "audit_result.go",
        "dao.go",
        "full.go",
        "mango.go",
        "modules.go",
        "panel.go",
        "pgc_cond.go",
        "playurl.go",
        "region.go",
        "sear_inter.go",
        "upbfs.go",
        "upper.go",
        "user.go",
    ],
    importpath = "go-common/app/admin/main/tv/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/tv/conf",
        "//app/admin/main/tv/model",
        "//library/cache/memcache",
        "//library/database/elastic",
        "//library/database/orm",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/time",
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
