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
        "local.go",
        "playurl.go",
        "search.go",
        "view.go",
    ],
    importpath = "go-common/app/interface/main/app-intl/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-card/model",
        "//app/interface/main/app-card/model/card",
        "//app/interface/main/app-card/model/card/ai",
        "//app/interface/main/app-intl/conf",
        "//app/interface/main/app-intl/model",
        "//app/interface/main/app-intl/model/feed",
        "//app/interface/main/app-intl/model/player",
        "//app/interface/main/app-intl/model/view",
        "//app/interface/main/app-intl/service/feed",
        "//app/interface/main/app-intl/service/player",
        "//app/interface/main/app-intl/service/search",
        "//app/interface/main/app-intl/service/view",
        "//library/conf/env",
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
