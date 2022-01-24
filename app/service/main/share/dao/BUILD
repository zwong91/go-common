package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "databus_test.go",
        "mysql_test.go",
        "redis_test.go",
        "share_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/share/conf",
        "//app/service/main/share/model",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "databus.go",
        "mysql.go",
        "redis.go",
        "share.go",
    ],
    importpath = "go-common/app/service/main/share/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/share/conf",
        "//app/service/main/share/model",
        "//library/cache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/ip",
        "//library/queue/databus",
        "//vendor/github.com/dgryski/go-farm",
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
