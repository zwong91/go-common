package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "archive_test",
    srcs = [
        "archive_test.go",
        "dao_test.go",
        "upper_rds_test.go",
        "upper_staff_test.go",
        "upper_test.go",
    ],
    embed = [":archive"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/up/conf",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "archive",
    srcs = [
        "archive.go",
        "dao.go",
        "upper.go",
        "upper_rds.go",
        "upper_staff.go",
    ],
    importpath = "go-common/app/service/main/up/dao/archive",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/up/conf",
        "//app/service/main/up/model",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/stat/prom",
        "//library/time",
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
