package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "http_test",
    srcs = [
        "academy_test.go",
        "whiteList_test.go",
    ],
    embed = [":http"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/creative/conf",
        "//app/admin/main/creative/service",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "http",
    srcs = [
        "academy_archive.go",
        "academy_occupation.go",
        "academy_tag.go",
        "app.go",
        "csv.go",
        "export.go",
        "http.go",
        "index.go",
        "local.go",
        "material.go",
        "material_category.go",
        "music.go",
        "music_category.go",
        "music_material.go",
        "music_with_category.go",
        "music_with_material.go",
        "operCollectArc.go",
        "task.go",
        "upload.go",
        "whitelist.go",
    ],
    importpath = "go-common/app/admin/main/creative/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/creative/conf",
        "//app/admin/main/creative/model/academy",
        "//app/admin/main/creative/model/app",
        "//app/admin/main/creative/model/logcli",
        "//app/admin/main/creative/model/material",
        "//app/admin/main/creative/model/music",
        "//app/admin/main/creative/model/operation",
        "//app/admin/main/creative/model/task",
        "//app/admin/main/creative/model/whitelist",
        "//app/admin/main/creative/service",
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/middleware/permit",
        "//library/net/http/blademaster/render",
        "//library/net/metadata",
        "//library/sync/errgroup",
        "//library/xstr",
        "//vendor/github.com/davecgh/go-spew/spew",
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
