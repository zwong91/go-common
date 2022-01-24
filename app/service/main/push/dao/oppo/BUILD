package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "oppo_test",
    srcs = ["client_test.go"],
    embed = [":oppo"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/push/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "oppo",
    srcs = [
        "auth.go",
        "client.go",
        "define.go",
    ],
    importpath = "go-common/app/service/main/push/dao/oppo",
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
