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
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-player/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "resource",
    srcs = ["dao.go"],
    importpath = "go-common/app/interface/main/app-player/dao/resource",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-player/conf",
        "//app/service/main/resource/rpc/client",
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
