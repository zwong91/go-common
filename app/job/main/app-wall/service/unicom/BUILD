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
        "//app/job/main/app-wall/conf",
        "//app/job/main/app-wall/model/unicom",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "unicom",
    srcs = [
        "infoc.go",
        "service.go",
        "unicom.go",
    ],
    importpath = "go-common/app/job/main/app-wall/service/unicom",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/app-wall/conf",
        "//app/job/main/app-wall/dao/seq",
        "//app/job/main/app-wall/dao/unicom",
        "//app/job/main/app-wall/model/unicom",
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
