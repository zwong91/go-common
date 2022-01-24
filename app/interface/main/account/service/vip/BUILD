package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "vip_test",
    srcs = [
        "associate_outer_test.go",
        "coupon_test.go",
        "manager_test.go",
        "panel_test.go",
        "vip_test.go",
    ],
    embed = [":vip"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/account/conf",
        "//app/interface/main/account/model",
        "//app/service/main/vip/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "vip",
    srcs = [
        "associate.go",
        "associate_limit.go",
        "associate_outer.go",
        "coupon.go",
        "manager.go",
        "panel.go",
        "privilege.go",
        "resource.go",
        "tips.go",
        "vip.go",
        "welfare.go",
    ],
    importpath = "go-common/app/interface/main/account/service/vip",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/account/conf",
        "//app/interface/main/account/dao/vip",
        "//app/interface/main/account/model",
        "//app/service/main/coupon/model",
        "//app/service/main/relation/model",
        "//app/service/main/relation/rpc/client",
        "//app/service/main/resource/model",
        "//app/service/main/resource/rpc/client",
        "//app/service/main/vip/api",
        "//app/service/main/vip/model",
        "//app/service/main/vip/rpc/client",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup",
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
