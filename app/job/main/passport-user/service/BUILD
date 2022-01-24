load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "aes.go",
        "aso_full_migration.go",
        "aso_incr_migration.go",
        "duplicate.go",
        "passport_util.go",
        "service.go",
        "user.go",
    ],
    importpath = "go-common/app/job/main/passport-user/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/passport-user/conf",
        "//app/job/main/passport-user/dao",
        "//app/job/main/passport-user/model",
        "//library/database/sql",
        "//library/log",
        "//library/queue/databus",
        "//library/queue/databus/databusutil",
        "//vendor/github.com/go-sql-driver/mysql",
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

go_test(
    name = "service_test",
    srcs = ["aes_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
)
