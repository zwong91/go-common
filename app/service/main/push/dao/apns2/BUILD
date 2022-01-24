package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "apns2_test",
    srcs = ["client_test.go"],
    embed = [":apns2"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/push/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "apns2",
    srcs = [
        "client.go",
        "notification.go",
    ],
    importpath = "go-common/app/service/main/push/dao/apns2",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//library/stat",
        "//library/stat/prom",
        "//vendor/golang.org/x/net/http2",
        "//vendor/golang.org/x/net/proxy",
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
