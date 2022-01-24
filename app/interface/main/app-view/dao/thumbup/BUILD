package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "thumbup_test",
    srcs = ["dao_test.go"],
    embed = [":thumbup"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-view/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "thumbup",
    srcs = ["dao.go"],
    importpath = "go-common/app/interface/main/app-view/dao/thumbup",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-view/conf",
        "//app/service/main/thumbup/api",
        "//library/net/metadata",
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
