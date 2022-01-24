package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "activity.go",
        "bp.go",
        "coupon.go",
        "http.go",
        "notify.go",
        "order.go",
        "order2.go",
        "panel.go",
        "point.go",
        "resouce.go",
        "sync.go",
        "tips.go",
        "user.go",
        "vip.go",
    ],
    importpath = "go-common/app/service/main/vip/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/vip/conf",
        "//app/service/main/vip/model",
        "//app/service/main/vip/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/metadata",
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
