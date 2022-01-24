package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "vip_test",
    srcs = [
        "cache_test.go",
        "dao_test.go",
        "mysql_test.go",
    ],
    embed = [":vip"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/xuser/conf",
        "//app/service/live/xuser/model",
        "//library/database/sql",
        "//library/log",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "vip",
    srcs = [
        "cache.go",
        "dao.go",
        "mysql.go",
    ],
    importpath = "go-common/app/service/live/xuser/dao/vip",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/xuser/conf",
        "//app/service/live/xuser/model",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/net/metadata",
        "//library/time",
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
