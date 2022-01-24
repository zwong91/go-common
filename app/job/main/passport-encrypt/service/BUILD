load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "aes.go",
        "aso.go",
        "aso_full_migration.go",
        "aso_incr_migration.go",
        "passport_util.go",
        "service.go",
    ],
    importpath = "go-common/app/job/main/passport-encrypt/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/passport-encrypt/conf",
        "//app/job/main/passport-encrypt/dao",
        "//app/job/main/passport-encrypt/model",
        "//library/log",
        "//library/queue/databus",
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
    srcs = [
        "aes_test.go",
        "aso_full_migration_test.go",
        "aso_incr_migration_test.go",
        "passport_util_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/passport-encrypt/conf",
        "//app/job/main/passport-encrypt/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
