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
        "//app/admin/main/push/conf",
        "//app/admin/main/push/model",
        "//app/service/main/push/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "dataplatform.go",
        "service.go",
        "task.go",
        "upload.go",
    ],
    importpath = "go-common/app/admin/main/push/service",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/push/conf",
        "//app/admin/main/push/dao",
        "//app/admin/main/push/model",
        "//app/service/main/push/model",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/sync/errgroup",
        "//vendor/github.com/jinzhu/gorm",
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
