load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "audit.go",
        "auth.go",
        "history.go",
        "http.go",
        "media.go",
        "others.go",
        "pages.go",
        "recommend.go",
        "search.go",
        "tvvip.go",
        "ugc.go",
        "ugc_view.go",
        "upgrade.go",
        "wild.go",
        "zone_index.go",
    ],
    importpath = "go-common/app/interface/main/tv/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/dao/thirdp",
        "//app/interface/main/tv/model",
        "//app/interface/main/tv/model/search",
        "//app/interface/main/tv/model/tvvip",
        "//app/interface/main/tv/service/app",
        "//app/interface/main/tv/service/audit",
        "//app/interface/main/tv/service/favorite",
        "//app/interface/main/tv/service/goblin",
        "//app/interface/main/tv/service/history",
        "//app/interface/main/tv/service/pgc",
        "//app/interface/main/tv/service/search",
        "//app/interface/main/tv/service/thirdp",
        "//app/interface/main/tv/service/tvvip",
        "//app/interface/main/tv/service/view",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/http/blademaster/render",
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
