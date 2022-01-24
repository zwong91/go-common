package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "banner.go",
        "daily.go",
        "http.go",
        "local.go",
        "rank.go",
        "region.go",
        "show.go",
    ],
    importpath = "go-common/app/interface/main/app-show/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-show/conf",
        "//app/interface/main/app-show/model",
        "//app/interface/main/app-show/model/region",
        "//app/interface/main/app-show/model/show",
        "//app/interface/main/app-show/service/banner",
        "//app/interface/main/app-show/service/daily",
        "//app/interface/main/app-show/service/ping",
        "//app/interface/main/app-show/service/rank",
        "//app/interface/main/app-show/service/region",
        "//app/interface/main/app-show/service/show",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
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
