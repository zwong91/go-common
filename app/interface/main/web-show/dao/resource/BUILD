package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "resource_test",
    srcs = ["dao_test.go"],
    embed = [":resource"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/web-show/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "resource",
    srcs = [
        "dao.go",
        "res.go",
        "vdoactive.go",
        "vdoad.go",
    ],
    importpath = "go-common/app/interface/main/web-show/dao/resource",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/web-show/conf",
        "//app/interface/main/web-show/model/resource",
        "//library/database/sql",
        "//library/log",
        "//library/stat/prom",
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
