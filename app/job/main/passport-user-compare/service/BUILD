package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "aes.go",
        "data_fix.go",
        "duplicate.go",
        "dynamic_compare.go",
        "full_compare.go",
        "service.go",
        "utils.go",
    ],
    importpath = "go-common/app/job/main/passport-user-compare/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/passport-user-compare/conf",
        "//app/job/main/passport-user-compare/dao",
        "//app/job/main/passport-user-compare/model",
        "//library/log",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/robfig/cron",
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
    name = "service_test",
    srcs = ["aes_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/passport-user-compare/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
