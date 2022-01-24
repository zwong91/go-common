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
        "rpc_test.go",
        "search_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/favorite/conf",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "fav.go",
        "rpc.go",
        "search.go",
        "service.go",
    ],
    importpath = "go-common/app/service/main/favorite/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/favorite/conf",
        "//app/service/main/favorite/dao",
        "//app/service/main/favorite/model",
        "//app/service/main/filter/model/rpc",
        "//app/service/main/filter/rpc/client",
        "//app/service/main/rank/api/gorpc",
        "//app/service/main/rank/model",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
        "//library/sync/errgroup",
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
