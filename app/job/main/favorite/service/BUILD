package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "fav_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/favorite/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "fav.go",
        "service.go",
        "stat.go",
    ],
    importpath = "go-common/app/job/main/favorite/service",
    tags = ["automanaged"],
    deps = [
        "//app/interface/openplatform/article/model",
        "//app/interface/openplatform/article/rpc/client",
        "//app/job/main/favorite/conf",
        "//app/job/main/favorite/dao/fav",
        "//app/job/main/favorite/dao/music",
        "//app/job/main/favorite/dao/pub",
        "//app/job/main/favorite/dao/stat",
        "//app/job/main/favorite/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/coin/api/gorpc",
        "//app/service/main/coin/model",
        "//app/service/main/favorite/model",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//library/sync/pipeline/fanout",
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
