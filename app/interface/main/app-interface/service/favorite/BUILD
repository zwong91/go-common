package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "favorite_test",
    srcs = [
        "favorite_test.go",
        "service_test.go",
    ],
    embed = [":favorite"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//app/interface/main/app-interface/model/favorite",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "favorite",
    srcs = [
        "favorite.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/app-interface/service/favorite",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//app/interface/main/app-interface/dao/article",
        "//app/interface/main/app-interface/dao/audio",
        "//app/interface/main/app-interface/dao/bangumi",
        "//app/interface/main/app-interface/dao/bplus",
        "//app/interface/main/app-interface/dao/favorite",
        "//app/interface/main/app-interface/dao/mall",
        "//app/interface/main/app-interface/dao/sp",
        "//app/interface/main/app-interface/dao/ticket",
        "//app/interface/main/app-interface/dao/topic",
        "//app/interface/main/app-interface/model",
        "//app/interface/main/app-interface/model/favorite",
        "//app/service/main/favorite/api",
        "//library/log",
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
