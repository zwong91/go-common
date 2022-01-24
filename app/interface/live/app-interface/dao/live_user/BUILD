package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "live_user",
    srcs = [
        "dahanghai.go",
        "dao.go",
        "get_tag.go",
    ],
    importpath = "go-common/app/interface/live/app-interface/dao/live_user",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/app-interface/conf",
        "//app/interface/live/app-interface/dao",
        "//app/service/live/live_user/api/liverpc/v1:liverpc",
        "//library/ecode",
        "//library/log",
        "//library/net/rpc/liverpc/context",
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
