package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "share_test",
    srcs = [
        "dao_test.go",
        "databus_test.go",
        "redis_test.go",
    ],
    embed = [":share"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "share",
    srcs = [
        "dao.go",
        "databus.go",
        "redis.go",
    ],
    importpath = "go-common/app/service/main/archive/dao/share",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/conf",
        "//library/cache/redis",
        "//library/log",
        "//library/net/ip",
        "//library/queue/databus",
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
