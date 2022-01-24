package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "feed.go",
        "http.go",
    ],
    importpath = "go-common/app/interface/main/web-feed/http",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/web-feed/conf",
        "//app/interface/main/web-feed/model",
        "//app/interface/main/web-feed/service",
        "//app/interface/openplatform/article/model",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/limit/aqm",
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
