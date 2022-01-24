package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "databus_test",
    srcs = [
        "archive_test.go",
        "dao_test.go",
    ],
    embed = [":databus"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/videoup/conf",
        "//app/admin/main/videoup/model/archive",
        "//app/admin/main/videoup/model/message",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "databus",
    srcs = [
        "dao.go",
        "redis.go",
    ],
    importpath = "go-common/app/admin/main/videoup/dao/databus",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/videoup/conf",
        "//app/admin/main/videoup/model/archive",
        "//app/admin/main/videoup/model/message",
        "//library/cache/redis",
        "//library/conf/env",
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
