package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "resource_test",
    srcs = ["res_test.go"],
    embed = [":resource"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/web-show/conf",
        "//app/interface/main/web-show/model/resource",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "resource",
    srcs = [
        "ad.go",
        "refresh.go",
        "res.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/web-show/service/resource",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/web-show/conf",
        "//app/interface/main/web-show/dao/ad",
        "//app/interface/main/web-show/dao/bangumi",
        "//app/interface/main/web-show/dao/data",
        "//app/interface/main/web-show/dao/resource",
        "//app/interface/main/web-show/model/resource",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/location/model",
        "//app/service/main/location/rpc/client",
        "//app/service/main/resource/rpc/client",
        "//library/log",
        "//library/net/metadata",
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
