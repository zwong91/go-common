package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "archive_test.go",
        "bangumi_test.go",
        "click_test.go",
        "dao_test.go",
        "forbid_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/click/conf",
        "//app/job/main/click/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "archive.go",
        "bangumi.go",
        "click.go",
        "dao.go",
        "forbid.go",
    ],
    importpath = "go-common/app/job/main/click/dao",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/click/conf",
        "//app/job/main/click/model",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
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
