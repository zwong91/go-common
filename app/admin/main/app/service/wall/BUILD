package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "wall_test",
    srcs = ["wall_test.go"],
    embed = [":wall"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/app/conf",
        "//app/admin/main/app/model/wall",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "wall",
    srcs = ["wall.go"],
    importpath = "go-common/app/admin/main/app/service/wall",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/app/conf",
        "//app/admin/main/app/dao/wall",
        "//app/admin/main/app/model/wall",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
