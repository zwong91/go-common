package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "manager_test",
    srcs = ["dao_test.go"],
    embed = [":manager"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/credit/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "manager",
    srcs = ["dao.go"],
    importpath = "go-common/app/admin/main/credit/dao/manager",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/credit/conf",
        "//app/admin/main/credit/model",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/sync/errgroup",
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
