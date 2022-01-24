package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "mysql_test",
    srcs = [
        "dao_test.go",
        "monitor_test.go",
        "task_test.go",
    ],
    embed = [":mysql"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/aegis/conf",
        "//app/admin/main/aegis/model/common",
        "//app/admin/main/aegis/model/task",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "mysql",
    srcs = [
        "dao.go",
        "monitor.go",
        "task.go",
    ],
    importpath = "go-common/app/admin/main/aegis/dao/mysql",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/aegis/conf",
        "//app/admin/main/aegis/model/common",
        "//app/admin/main/aegis/model/monitor",
        "//app/admin/main/aegis/model/task",
        "//library/database/sql",
        "//library/log",
        "//library/xstr",
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
