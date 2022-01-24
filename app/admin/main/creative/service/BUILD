package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "service_test.go",
        "upbfs_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/creative/conf",
        "//app/service/main/account/api",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "academy.go",
        "log.go",
        "material.go",
        "music.go",
        "service.go",
        "task.go",
        "upbfs.go",
        "whitelist.go",
    ],
    importpath = "go-common/app/admin/main/creative/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/creative/conf",
        "//app/admin/main/creative/dao",
        "//app/admin/main/creative/model/academy",
        "//app/admin/main/creative/model/bfs",
        "//app/admin/main/creative/model/logcli",
        "//app/admin/main/creative/model/material",
        "//app/admin/main/creative/model/music",
        "//app/admin/main/creative/model/task",
        "//app/admin/main/creative/model/whitelist",
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//library/conf/env",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/queue/databus/report",
        "//library/xstr",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/golang.org/x/sync/errgroup",
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
