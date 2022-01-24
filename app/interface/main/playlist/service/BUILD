package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "favorite_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/playlist/conf",
        "//app/interface/main/playlist/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "favorite.go",
        "service.go",
        "stat.go",
        "video.go",
    ],
    importpath = "go-common/app/interface/main/playlist/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/playlist/conf",
        "//app/interface/main/playlist/dao",
        "//app/interface/main/playlist/model",
        "//app/service/main/account/api",
        "//app/service/main/account/rpc/client",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/favorite/api/gorpc",
        "//app/service/main/favorite/model",
        "//app/service/main/filter/rpc/client",
        "//library/cache",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup",
        "//library/time",
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
