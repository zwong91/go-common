package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "redis_test",
    srcs = [
        "dao_test.go",
        "monitor_test.go",
        "net_test.go",
        "task_test.go",
    ],
    embed = [":redis"],
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
    name = "redis",
    srcs = [
        "cluster.go",
        "dao.go",
        "monitor.go",
        "net.go",
        "task.go",
    ],
    importpath = "go-common/app/admin/main/aegis/dao/redis",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/aegis/conf",
        "//app/admin/main/aegis/model/common",
        "//app/admin/main/aegis/model/monitor",
        "//app/admin/main/aegis/model/task",
        "//library/cache/redis",
        "//library/ecode",
        "//library/log",
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
