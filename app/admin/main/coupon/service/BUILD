package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "activity.go",
        "allowance.go",
        "code.go",
        "coupon.go",
        "service.go",
        "view.go",
    ],
    importpath = "go-common/app/admin/main/coupon/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/coupon/conf",
        "//app/admin/main/coupon/dao",
        "//app/admin/main/coupon/model",
        "//app/service/main/coupon/model",
        "//app/service/main/coupon/rpc/client",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/pipeline/fanout",
        "//library/time",
        "//library/xstr",
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

go_test(
    name = "service_test",
    srcs = [
        "activity_test.go",
        "allowance_test.go",
        "code_test.go",
        "coupon_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/coupon/conf",
        "//app/admin/main/coupon/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
