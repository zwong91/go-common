package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "oplog_test",
    srcs = [
        "dao_test.go",
        "query_test.go",
    ],
    embed = [":oplog"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/dm/conf",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "oplog",
    srcs = [
        "dao.go",
        "query.go",
    ],
    importpath = "go-common/app/admin/main/dm/dao/oplog",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/dm/conf",
        "//app/admin/main/dm/model/oplog",
        "//library/conf/env",
        "//library/log",
        "//library/net/http/blademaster",
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
