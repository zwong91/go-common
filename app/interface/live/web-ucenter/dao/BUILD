package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = ["dao.go"],
    importpath = "go-common/app/interface/live/web-ucenter/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/web-ucenter/conf",
        "//app/service/live/rankdb/api/liverpc",
        "//app/service/live/rc/api/liverpc",
        "//app/service/live/room/api/liverpc",
        "//library/net/rpc/liverpc",
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
        "//app/interface/live/web-ucenter/dao/capsule:all-srcs",
        "//app/interface/live/web-ucenter/dao/history:all-srcs",
        "//app/interface/live/web-ucenter/dao/user:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
