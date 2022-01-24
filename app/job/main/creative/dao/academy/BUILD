package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "academy",
    srcs = [
        "academy.go",
        "dao.go",
    ],
    importpath = "go-common/app/job/main/creative/dao/academy",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/creative/conf",
        "//app/job/main/creative/model",
        "//library/database/sql",
        "//library/log",
        "//library/xstr",
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
    name = "academy_test",
    srcs = [
        "academy_test.go",
        "dao_test.go",
    ],
    embed = [":academy"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/creative/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
