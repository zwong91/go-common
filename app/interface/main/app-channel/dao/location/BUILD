package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "location_test",
    srcs = ["dao_test.go"],
    embed = [":location"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-channel/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "location",
    srcs = ["dao.go"],
    importpath = "go-common/app/interface/main/app-channel/dao/location",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-channel/conf",
        "//app/service/main/location/model",
        "//app/service/main/location/rpc/client",
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
