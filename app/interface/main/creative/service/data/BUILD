package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "data_test",
    srcs = ["service_test.go"],
    embed = [":data"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/service",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "data",
    srcs = [
        "app.go",
        "archive.go",
        "article.go",
        "fan.go",
        "overview.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/creative/service/data",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/account",
        "//app/interface/main/creative/dao/activity",
        "//app/interface/main/creative/dao/archive",
        "//app/interface/main/creative/dao/article",
        "//app/interface/main/creative/dao/bfs",
        "//app/interface/main/creative/dao/data",
        "//app/interface/main/creative/dao/medal",
        "//app/interface/main/creative/dao/tag",
        "//app/interface/main/creative/model/archive",
        "//app/interface/main/creative/model/data",
        "//app/interface/main/creative/model/medal",
        "//app/interface/main/creative/model/tag",
        "//app/interface/main/creative/service",
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//library/ecode",
        "//library/log",
        "//library/stat/prom",
        "//library/sync/errgroup",
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
