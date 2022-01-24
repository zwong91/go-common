package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "account.go",
        "album.go",
        "app.go",
        "archive.go",
        "article.go",
        "assist.go",
        "bangumi.go",
        "blacklist.go",
        "channel.go",
        "coin.go",
        "dynamic.go",
        "fav.go",
        "http.go",
        "setting.go",
        "shop.go",
        "tag.go",
        "web.go",
    ],
    importpath = "go-common/app/interface/main/space/http",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/space/conf",
        "//app/interface/main/space/model",
        "//app/interface/main/space/service",
        "//app/interface/openplatform/article/model",
        "//app/service/main/assist/model/assist",
        "//library/conf/env",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/supervisor",
        "//library/net/http/blademaster/middleware/verify",
        "//library/xstr",
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
