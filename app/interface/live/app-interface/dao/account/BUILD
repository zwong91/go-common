package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "account",
    srcs = [
        "dao.go",
        "userInfo.go",
    ],
    importpath = "go-common/app/interface/live/app-interface/dao/account",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/app-interface/conf",
        "//app/service/main/account/api",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
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
