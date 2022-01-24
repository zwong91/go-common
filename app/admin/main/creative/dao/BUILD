package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "academy_test.go",
        "dao_test.go",
        "material_test.go",
        "msg_test.go",
        "search_test.go",
        "upbfs_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/creative/conf",
        "//app/admin/main/creative/model/academy",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "academy.go",
        "dao.go",
        "material.go",
        "msg.go",
        "search.go",
        "upbfs.go",
    ],
    importpath = "go-common/app/admin/main/creative/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/creative/conf",
        "//app/admin/main/creative/model/academy",
        "//app/admin/main/creative/model/material",
        "//app/interface/openplatform/article/model",
        "//app/interface/openplatform/article/rpc/client",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//library/database/elastic",
        "//library/database/orm",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/xstr",
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
