package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "bnj",
    srcs = [
        "bnj.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/activity/service/bnj",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/dao/bnj",
        "//app/interface/main/activity/dao/like",
        "//app/interface/main/activity/model/bnj",
        "//app/service/main/archive/api",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
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
