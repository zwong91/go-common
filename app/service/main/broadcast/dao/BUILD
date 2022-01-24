package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "databus.go",
        "redis.go",
    ],
    importpath = "go-common/app/service/main/broadcast/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/broadcast/model",
        "//library/cache/redis",
        "//library/conf/paladin",
        "//library/log",
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

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "databus_test.go",
        "redis_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/broadcast/model",
        "//library/conf/paladin",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
