package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "relation",
    srcs = [
        "AppRelationFilterStrategy.go",
        "AppRelationSortStrategy.go",
        "Tools.go",
        "rpcWraper.go",
    ],
    importpath = "go-common/app/interface/live/app-interface/service/v1/relation",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/app-interface/api/http/v1:http",
        "//app/interface/live/app-interface/conf",
        "//app/interface/live/app-interface/dao",
        "//app/service/live/av/api/liverpc/v1:liverpc",
        "//app/service/live/fans_medal/api/liverpc/v1:liverpc",
        "//app/service/live/live_data/api/liverpc/v1:liverpc",
        "//app/service/live/relation/api/liverpc/v1:liverpc",
        "//app/service/live/room/api/liverpc/v1:liverpc",
        "//app/service/live/room/api/liverpc/v2:liverpc",
        "//app/service/live/room_ex/api/liverpc/v1:liverpc",
        "//app/service/live/userext/api/liverpc/v1:liverpc",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/net/rpc/liverpc/context",
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
