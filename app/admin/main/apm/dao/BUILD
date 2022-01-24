package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "canal_test.go",
        "dao_test.go",
        "ecode_test.go",
        "need_test.go",
        "pprof_test.go",
        "upload_test.go",
        "ut_app_test.go",
        "ut_rank_test.go",
        "ut_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/apm/conf",
        "//app/admin/main/apm/model/canal",
        "//app/admin/main/apm/model/need",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "canal.go",
        "dao.go",
        "ecode.go",
        "monkey.go",
        "need.go",
        "pprof.go",
        "upload.go",
        "ut.go",
        "ut_app.go",
        "ut_rank.go",
    ],
    importpath = "go-common/app/admin/main/apm/dao",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/apm/conf",
        "//app/admin/main/apm/model/canal",
        "//app/admin/main/apm/model/ecode",
        "//app/admin/main/apm/model/need",
        "//app/admin/main/apm/model/pprof",
        "//app/admin/main/apm/model/ut",
        "//library/cache/redis",
        "//library/conf/env",
        "//library/database/orm",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/github.com/bouk/monkey",
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
    srcs = [
        ":package-srcs",
        "//app/admin/main/apm/dao/mock:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
