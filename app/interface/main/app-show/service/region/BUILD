package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "region_test",
    srcs = [
        "region_test.go",
        "show_test.go",
    ],
    embed = [":region"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-show/conf",
        "//app/interface/main/app-show/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "region",
    srcs = [
        "audit.go",
        "banner.go",
        "cache.go",
        "infoc.go",
        "region.go",
        "service.go",
        "show.go",
    ],
    importpath = "go-common/app/interface/main/app-show/service/region",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-show/conf",
        "//app/interface/main/app-show/dao/activity",
        "//app/interface/main/app-show/dao/ad",
        "//app/interface/main/app-show/dao/archive",
        "//app/interface/main/app-show/dao/audit",
        "//app/interface/main/app-show/dao/bangumi",
        "//app/interface/main/app-show/dao/card",
        "//app/interface/main/app-show/dao/dynamic",
        "//app/interface/main/app-show/dao/location",
        "//app/interface/main/app-show/dao/recommend",
        "//app/interface/main/app-show/dao/region",
        "//app/interface/main/app-show/dao/resource",
        "//app/interface/main/app-show/dao/search",
        "//app/interface/main/app-show/dao/tag",
        "//app/interface/main/app-show/model",
        "//app/interface/main/app-show/model/bangumi",
        "//app/interface/main/app-show/model/banner",
        "//app/interface/main/app-show/model/card",
        "//app/interface/main/app-show/model/region",
        "//app/interface/main/app-show/model/tag",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//app/service/main/location/model",
        "//app/service/main/resource/model",
        "//app/service/openplatform/pgc-season/api/grpc/season/v1:season",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//library/net/metadata",
        "//library/stat/prom",
        "//library/time",
        "//vendor/github.com/dgryski/go-farm",
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
