package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "relation",
    srcs = [
        "BaseInfo.go",
        "Feed.go",
        "dao.go",
    ],
    importpath = "go-common/app/interface/live/app-interface/dao/relation",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/app-interface/conf",
        "//app/interface/live/app-interface/dao",
        "//app/service/live/relation/api/liverpc/v1:liverpc",
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
