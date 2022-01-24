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
        "api.go",
        "cache.go",
        "dao.go",
    ],
    importpath = "go-common/app/job/main/app-wall/dao/unicom",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/app-wall/conf",
        "//app/job/main/app-wall/model",
        "//app/job/main/app-wall/model/unicom",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
