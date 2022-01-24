package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "unicom_test",
    srcs = ["unicom_test.go"],
    embed = [":unicom"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-wall/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "unicom",
    srcs = [
        "cache.go",
        "databus.go",
        "infoc.go",
        "pack.go",
        "service.go",
        "unicom.go",
    ],
    importpath = "go-common/app/interface/main/app-wall/service/unicom",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-wall/conf",
        "//app/interface/main/app-wall/dao/account",
        "//app/interface/main/app-wall/dao/live",
        "//app/interface/main/app-wall/dao/seq",
        "//app/interface/main/app-wall/dao/shopping",
        "//app/interface/main/app-wall/dao/unicom",
        "//app/interface/main/app-wall/model",
        "//app/interface/main/app-wall/model/unicom",
        "//app/service/main/account/model",
        "//library/cache/memcache",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//library/queue/databus",
        "//library/queue/databus/report",
        "//library/stat/prom",
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
