package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "anniversary.go",
        "apply.go",
        "archive.go",
        "article.go",
        "author.go",
        "cards.go",
        "categories.go",
        "complaint.go",
        "creative.go",
        "favorite.go",
        "hotspots.go",
        "http.go",
        "like.go",
        "list.go",
        "notice.go",
        "rank.go",
        "read.go",
        "recommends.go",
        "search.go",
        "sentinel.go",
        "share.go",
        "users.go",
        "view.go",
    ],
    importpath = "go-common/app/interface/openplatform/article/http",
    tags = ["automanaged"],
    deps = [
        "//app/interface/openplatform/article/conf",
        "//app/interface/openplatform/article/dao",
        "//app/interface/openplatform/article/model",
        "//app/interface/openplatform/article/service",
        "//app/service/main/archive/api",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/antispam",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/cache",
        "//library/net/http/blademaster/middleware/cache/store",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/metadata",
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
