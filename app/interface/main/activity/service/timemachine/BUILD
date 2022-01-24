package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "timemachine",
    srcs = [
        "service.go",
        "timemachine.go",
    ],
    importpath = "go-common/app/interface/main/activity/service/timemachine",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/dao/like",
        "//app/interface/main/activity/dao/timemachine",
        "//app/interface/main/activity/model/timemachine",
        "//app/interface/main/tag/api",
        "//app/interface/openplatform/article/model",
        "//app/interface/openplatform/article/rpc/client",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup.v2:errgroup_v2",
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
