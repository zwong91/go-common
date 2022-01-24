package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "newcomer_test",
    srcs = ["service_test.go"],
    embed = [":newcomer"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/service",
    ],
)

go_library(
    name = "newcomer",
    srcs = [
        "h5.go",
        "service.go",
        "state.go",
        "task.go",
    ],
    importpath = "go-common/app/interface/main/creative/service/newcomer",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/academy",
        "//app/interface/main/creative/dao/account",
        "//app/interface/main/creative/dao/archive",
        "//app/interface/main/creative/dao/article",
        "//app/interface/main/creative/dao/data",
        "//app/interface/main/creative/dao/medal",
        "//app/interface/main/creative/dao/newcomer",
        "//app/interface/main/creative/dao/order",
        "//app/interface/main/creative/dao/watermark",
        "//app/interface/main/creative/model/newcomer",
        "//app/interface/main/creative/service",
        "//app/service/main/account/api",
        "//library/conf/env",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/queue/databus/report",
        "//library/sync/errgroup.v2:errgroup_v2",
        "//library/time",
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
