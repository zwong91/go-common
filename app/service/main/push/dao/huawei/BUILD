package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "huawei_test",
    srcs = [
        "access_test.go",
        "client_test.go",
    ],
    embed = [":huawei"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/smartystreets/goconvey/convey"],
)

go_library(
    name = "huawei",
    srcs = [
        "access.go",
        "client.go",
        "message.go",
    ],
    importpath = "go-common/app/service/main/push/dao/huawei",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//library/stat",
        "//library/stat/prom",
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
