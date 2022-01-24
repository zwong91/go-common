package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "info.go",
        "service.go",
        "utils.go",
    ],
    importpath = "go-common/app/service/main/vipinfo/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/vipinfo/api",
        "//app/service/main/vipinfo/conf",
        "//app/service/main/vipinfo/dao",
        "//app/service/main/vipinfo/model",
        "//library/ecode",
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

go_test(
    name = "service_test",
    srcs = ["info_test.go"],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/vipinfo/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
