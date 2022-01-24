package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = ["http.go"],
    importpath = "go-common/app/interface/live/app-interface/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/app-interface/api/http/v1:http",
        "//app/interface/live/app-interface/api/http/v2:http",
        "//app/interface/live/app-interface/conf",
        "//app/interface/live/app-interface/dao",
        "//app/interface/live/app-interface/service",
        "//app/interface/live/app-interface/service/v1:service",
        "//app/interface/live/app-interface/service/v1/app_conf",
        "//app/interface/live/app-interface/service/v2:service",
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
