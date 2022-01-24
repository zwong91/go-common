package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "allowance_test.go",
        "code_test.go",
        "coupon_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/coupon/api",
        "//app/service/main/coupon/conf",
        "//app/service/main/coupon/model",
        "//library/database/sql",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "allowance.go",
        "branch.go",
        "code.go",
        "coupon.go",
        "coupon_cartoon.go",
        "service.go",
    ],
    importpath = "go-common/app/service/main/coupon/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/coupon/api",
        "//app/service/main/coupon/conf",
        "//app/service/main/coupon/dao",
        "//app/service/main/coupon/model",
        "//app/service/main/vipinfo/api",
        "//app/service/main/vipinfo/model",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/sync/pipeline/fanout",
        "//library/time",
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
