package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = ["http.go"],
    importpath = "go-common/app/interface/live/live-demo/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/live-demo/api/http",
        "//app/interface/live/live-demo/api/http/v2:http",
        "//app/interface/live/live-demo/conf",
        "//app/interface/live/live-demo/dao",
        "//app/interface/live/live-demo/service",
        "//app/interface/live/live-demo/service/v2:service",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/verify",
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
