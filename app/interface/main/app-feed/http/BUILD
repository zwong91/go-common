package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "cache.go",
        "dynamic.go",
        "feed.go",
        "http.go",
        "local.go",
        "region.go",
    ],
    importpath = "go-common/app/interface/main/app-feed/http",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-card/model",
        "//app/interface/main/app-card/model/card",
        "//app/interface/main/app-card/model/card/ai",
        "//app/interface/main/app-card/model/card/operate",
        "//app/interface/main/app-feed/conf",
        "//app/interface/main/app-feed/model",
        "//app/interface/main/app-feed/model/feed",
        "//app/interface/main/app-feed/service/external",
        "//app/interface/main/app-feed/service/feed",
        "//app/interface/main/app-feed/service/ping",
        "//app/interface/main/app-feed/service/region",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/metadata",
        "//library/stat/prom",
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
