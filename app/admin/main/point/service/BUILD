package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = ["service_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/point/conf",
        "//app/admin/main/point/model",
        "//app/service/main/point/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "point.go",
        "service.go",
    ],
    importpath = "go-common/app/admin/main/point/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/point/conf",
        "//app/admin/main/point/dao",
        "//app/admin/main/point/model",
        "//app/service/main/point/model",
        "//app/service/main/point/rpc/client",
        "//library/ecode",
        "//library/log",
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
