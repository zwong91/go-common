package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "dialog_test.go",
        "jointly_test.go",
        "order_test.go",
        "panel_test.go",
        "privilege_test.go",
        "push_test.go",
        "resouce_test.go",
        "service_test.go",
        "tips_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/vip/conf",
        "//app/admin/main/vip/model",
        "//library/ecode",
        "//library/log",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "business.go",
        "dialog.go",
        "jointly.go",
        "month.go",
        "order.go",
        "panel.go",
        "platform.go",
        "privilege.go",
        "push.go",
        "resouce.go",
        "service.go",
        "tips.go",
        "version.go",
        "vip.go",
        "welfare.go",
    ],
    importpath = "go-common/app/admin/main/vip/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/vip/conf",
        "//app/admin/main/vip/dao",
        "//app/admin/main/vip/model",
        "//app/service/main/vip/model",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/time",
        "//vendor/github.com/google/uuid",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/pkg/errors",
        "//vendor/golang.org/x/sync/errgroup",
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
