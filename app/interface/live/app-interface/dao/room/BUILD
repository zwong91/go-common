package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "room",
    srcs = [
        "app.go",
        "appIndex.go",
        "dao.go",
        "info.go",
        "pendant.go",
        "recommend.go",
        "status.go",
        "uid2roomid.go",
    ],
    importpath = "go-common/app/interface/live/app-interface/dao/room",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/app-interface/conf",
        "//app/interface/live/app-interface/dao",
        "//app/service/live/room/api/liverpc/v1:liverpc",
        "//app/service/live/room/api/liverpc/v2:liverpc",
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
