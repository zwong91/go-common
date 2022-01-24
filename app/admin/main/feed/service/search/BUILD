package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "search",
    srcs = [
        "search.go",
        "search_web.go",
    ],
    importpath = "go-common/app/admin/main/feed/service/search",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/feed/conf",
        "//app/admin/main/feed/dao/search",
        "//app/admin/main/feed/dao/show",
        "//app/admin/main/feed/model/common",
        "//app/admin/main/feed/model/search",
        "//app/admin/main/feed/model/show",
        "//app/admin/main/feed/util",
        "//library/cache/memcache",
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/github.com/jinzhu/gorm",
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
    name = "search_test",
    srcs = ["search_test.go"],
    embed = [":search"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/feed/conf",
        "//app/admin/main/feed/model/search",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
