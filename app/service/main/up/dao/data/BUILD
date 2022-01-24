package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "data_test",
    srcs = [
        "dao_test.go",
        "field_test.go",
        "hbase_test.go",
    ],
    embed = [":data"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/up/conf",
        "//app/service/main/up/dao",
        "//library/database/hbase.v2:hbase_v2",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/github.com/tsuna/gohbase/hrpc",
    ],
)

go_library(
    name = "data",
    srcs = [
        "dao.go",
        "field.go",
        "hbase.go",
    ],
    importpath = "go-common/app/service/main/up/dao/data",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/up/conf",
        "//app/service/main/up/model/data",
        "//library/database/hbase.v2:hbase_v2",
        "//library/ecode",
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
