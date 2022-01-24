package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = ["service.go"],
    importpath = "go-common/app/service/live/xuser/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/xuser/conf",
        "//app/service/live/xuser/service/exp/v1:exp",
        "//app/service/live/xuser/service/guard/v1:guard",
        "//app/service/live/xuser/service/roomAdmin/v1:roomAdmin",
        "//app/service/live/xuser/service/vip/v1:vip",
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
    srcs = [
        ":package-srcs",
        "//app/service/live/xuser/service/exp/v1:all-srcs",
        "//app/service/live/xuser/service/guard/v1:all-srcs",
        "//app/service/live/xuser/service/roomAdmin/v1:all-srcs",
        "//app/service/live/xuser/service/vip/v1:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
