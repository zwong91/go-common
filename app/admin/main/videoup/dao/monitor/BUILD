package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "monitor_test",
    srcs = [
        "dao_test.go",
        "redis_test.go",
    ],
    embed = [":monitor"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/videoup/conf",
        "//app/admin/main/videoup/model/monitor",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "monitor",
    srcs = [
        "dao.go",
        "redis.go",
    ],
    importpath = "go-common/app/admin/main/videoup/dao/monitor",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/videoup/conf",
        "//app/admin/main/videoup/model/monitor",
        "//library/cache/redis",
        "//library/log",
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
