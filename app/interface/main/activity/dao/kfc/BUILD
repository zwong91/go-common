package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "kfc",
    srcs = [
        "cache.go",
        "dao.cache.go",
        "dao.go",
        "kfc.go",
        "mc.cache.go",
    ],
    importpath = "go-common/app/interface/main/activity/dao/kfc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/model/kfc",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/stat/prom",
        "//vendor/github.com/pkg/errors",
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
    name = "kfc_test",
    srcs = [
        "cache_test.go",
        "dao.cache_test.go",
        "dao_test.go",
        "kfc_test.go",
        "mc.cache_test.go",
    ],
    embed = [":kfc"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/model/kfc",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
