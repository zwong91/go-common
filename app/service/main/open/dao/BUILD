package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "app.go",
        "dao.go",
        "memcache.go",
    ],
    importpath = "go-common/app/service/main/open/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/open/conf",
        "//library/cache/memcache",
        "//library/database/orm",
        "//library/ecode",
        "//library/log",
        "//vendor/github.com/jinzhu/gorm",
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
    name = "dao_test",
    srcs = [
        "app_test.go",
        "dao_test.go",
        "memcache_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/open/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
