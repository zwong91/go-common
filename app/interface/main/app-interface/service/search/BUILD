package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "search_test",
    srcs = ["search_test.go"],
    embed = [":search"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "search",
    srcs = ["search.go"],
    importpath = "go-common/app/interface/main/app-interface/service/search",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//app/interface/main/app-interface/dao/account",
        "//app/interface/main/app-interface/dao/archive",
        "//app/interface/main/app-interface/dao/article",
        "//app/interface/main/app-interface/dao/bangumi",
        "//app/interface/main/app-interface/dao/bplus",
        "//app/interface/main/app-interface/dao/live",
        "//app/interface/main/app-interface/dao/resource",
        "//app/interface/main/app-interface/dao/search",
        "//app/interface/main/app-interface/dao/tag",
        "//app/interface/main/app-interface/model",
        "//app/interface/main/app-interface/model/bangumi",
        "//app/interface/main/app-interface/model/banner",
        "//app/interface/main/app-interface/model/bplus",
        "//app/interface/main/app-interface/model/live",
        "//app/interface/main/app-interface/model/search",
        "//app/interface/main/app-interface/model/tag",
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/resource/model",
        "//app/service/openplatform/pgc-season/api/grpc/season/v1:season",
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
