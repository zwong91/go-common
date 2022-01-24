package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "esports",
    srcs = ["service.go"],
    importpath = "go-common/app/job/main/web-goblin/service/esports",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/web-goblin/conf",
        "//app/job/main/web-goblin/dao/esports",
        "//app/job/main/web-goblin/model/esports",
        "//app/service/main/archive/api",
        "//app/service/main/favorite/api/gorpc",
        "//app/service/main/favorite/model",
        "//library/ecode",
        "//library/log",
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
