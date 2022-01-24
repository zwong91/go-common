package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "archive_stat_test.go",
        "csv_test.go",
        "service_test.go",
        "task_log_test.go",
        "task_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/laser/conf",
        "//app/admin/main/laser/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "archive_stat.go",
        "csv.go",
        "service.go",
        "task.go",
        "task_log.go",
    ],
    importpath = "go-common/app/admin/main/laser/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/laser/conf",
        "//app/admin/main/laser/dao",
        "//app/admin/main/laser/model",
        "//library/log",
        "//library/time",
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
