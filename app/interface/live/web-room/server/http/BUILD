package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "dm.go",
        "http.go",
    ],
    importpath = "go-common/app/interface/live/web-room/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/web-room/api/http/v1:http",
        "//app/interface/live/web-room/conf",
        "//app/interface/live/web-room/service",
        "//app/interface/live/web-room/service/v1:service",
        "//app/interface/live/web-room/service/v1/dm",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
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
