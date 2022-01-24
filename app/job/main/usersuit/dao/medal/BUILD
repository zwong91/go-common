package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "medal_test",
    srcs = ["dao_test.go"],
    embed = [":medal"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/usersuit/conf",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "medal",
    srcs = [
        "dao.go",
        "data.go",
        "rpc.go",
    ],
    importpath = "go-common/app/job/main/usersuit/dao/medal",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/usersuit/conf",
        "//app/job/main/usersuit/model",
        "//app/service/main/usersuit/model",
        "//app/service/main/usersuit/rpc/client",
        "//library/cache/memcache",
        "//library/database/sql",
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
