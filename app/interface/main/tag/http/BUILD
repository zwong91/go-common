package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "action.go",
        "arctag.go",
        "change.go",
        "channel.go",
        "http.go",
        "platrform.go",
        "sub.go",
        "tag.go",
        "tagarc.go",
    ],
    importpath = "go-common/app/interface/main/tag/http",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tag/conf",
        "//app/interface/main/tag/model",
        "//app/interface/main/tag/service",
        "//app/service/main/archive/api",
        "//app/service/main/tag/model",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
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
