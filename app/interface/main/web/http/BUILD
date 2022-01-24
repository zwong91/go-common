package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "account.go",
        "archive.go",
        "article.go",
        "bnj.go",
        "broadcast.go",
        "coin.go",
        "dynamic.go",
        "elec.go",
        "feedback.go",
        "help.go",
        "http.go",
        "icon.go",
        "ip.go",
        "like.go",
        "nav.go",
        "newlist.go",
        "online.go",
        "ranking.go",
        "resource.go",
        "search.go",
        "tag.go",
        "wechat.go",
    ],
    importpath = "go-common/app/interface/main/web/http",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/web/conf",
        "//app/interface/main/web/model",
        "//app/interface/main/web/service",
        "//app/interface/openplatform/article/model",
        "//app/service/main/archive/api",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/cache",
        "//library/net/http/blademaster/middleware/cache/store",
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
