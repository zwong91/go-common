package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "pay",
    srcs = ["service.go"],
    importpath = "go-common/app/interface/main/creative/service/pay",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/faq",
        "//app/interface/main/creative/dao/pay",
        "//app/interface/main/creative/dao/up",
        "//app/interface/main/creative/model/faq",
        "//app/interface/main/creative/service",
        "//library/log",
        "//library/sync/pipeline/fanout",
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
