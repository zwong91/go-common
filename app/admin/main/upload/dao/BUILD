package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "bfs_test.go",
        "dao_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/upload/conf",
        "//app/admin/main/upload/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "bfs.go",
        "dao.go",
        "hbase.go",
    ],
    importpath = "go-common/app/admin/main/upload/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/upload/conf",
        "//app/admin/main/upload/model",
        "//library/ecode",
        "//library/log",
        "//vendor/github.com/tsuna/gohbase",
        "//vendor/github.com/tsuna/gohbase/hrpc",
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
