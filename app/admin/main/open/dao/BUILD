package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "app_test.go",
        "dao_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/open/conf",
        "//app/admin/main/open/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "app.go",
        "dao.go",
    ],
    importpath = "go-common/app/admin/main/open/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/open/conf",
        "//app/admin/main/open/model",
        "//library/database/orm",
        "//library/log",
        "//vendor/github.com/jinzhu/gorm",
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
