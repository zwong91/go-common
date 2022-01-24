package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "ads_test.go",
        "banner_test.go",
        "relate_test.go",
        "resource_test.go",
        "service_test.go",
        "special_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/resource/api/v1:api",
        "//app/service/main/resource/conf",
        "//library/cache/redis",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "abtest.go",
        "ads.go",
        "archive.go",
        "banner.go",
        "relate.go",
        "resource.go",
        "service.go",
        "show.go",
        "special.go",
    ],
    importpath = "go-common/app/service/main/resource/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/location/model",
        "//app/service/main/location/rpc/client",
        "//app/service/main/resource/api/v1:api",
        "//app/service/main/resource/conf",
        "//app/service/main/resource/dao/abtest",
        "//app/service/main/resource/dao/ads",
        "//app/service/main/resource/dao/alarm",
        "//app/service/main/resource/dao/cpm",
        "//app/service/main/resource/dao/manager",
        "//app/service/main/resource/dao/resource",
        "//app/service/main/resource/dao/show",
        "//app/service/main/resource/model",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//library/sync/errgroup",
        "//library/xstr",
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
