package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "mc_test",
    srcs = [
        "dao_test.go",
        "task_test.go",
    ],
    embed = [":mc"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/aegis/conf",
        "//app/admin/main/aegis/model/common",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "mc",
    srcs = [
        "dao.go",
        "task.go",
    ],
    importpath = "go-common/app/admin/main/aegis/dao/mc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/aegis/conf",
        "//app/admin/main/aegis/model/common",
        "//library/cache/memcache",
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
