package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "favorite.go",
        "match.go",
        "match_active.go",
        "pointdata.go",
        "search.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/esports/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/esports/conf",
        "//app/interface/main/esports/dao",
        "//app/interface/main/esports/model",
        "//app/service/main/archive/api",
        "//app/service/main/favorite/api/gorpc",
        "//app/service/main/favorite/model",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
        "//library/xstr",
        "//vendor/github.com/robfig/cron",
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

go_test(
    name = "service_test",
    srcs = [
        "favorite_test.go",
        "match_test.go",
        "search_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/esports/conf",
        "//app/interface/main/esports/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
