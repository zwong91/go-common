package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "roomAdmin.go",
    ],
    importpath = "go-common/app/interface/live/web-room/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/web-room/api/http/v1:http",
        "//app/interface/live/web-room/conf",
        "//app/service/live/xuser/api/grpc/v1:grpc",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//library/log",
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
