package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "upload_test.go",
        "utils_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/up/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "upload.go",
        "utils.go",
    ],
    importpath = "go-common/app/admin/main/up/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up/conf",
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
    srcs = [
        ":package-srcs",
        "//app/admin/main/up/dao/data:all-srcs",
        "//app/admin/main/up/dao/email:all-srcs",
        "//app/admin/main/up/dao/global:all-srcs",
        "//app/admin/main/up/dao/manager:all-srcs",
        "//app/admin/main/up/dao/tag:all-srcs",
        "//app/admin/main/up/dao/upcrm:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
