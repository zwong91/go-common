package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "dm.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/live/open-interface/internal/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/open-interface/api/http/v1:http",
        "//app/interface/live/open-interface/internal/dao",
        "//app/service/live/broadcast-proxy/api/v1:api",
        "//app/service/live/resource/sdk",
        "//library/conf/paladin",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
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
