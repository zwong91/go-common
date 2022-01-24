package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "business.go",
        "channel.go",
        "channel_category.go",
        "channel_rule.go",
        "http.go",
        "limit.go",
        "rank.go",
        "relation.go",
        "report.go",
        "resource.go",
        "synonym.go",
        "tag.go",
    ],
    importpath = "go-common/app/admin/main/tag/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/tag/conf",
        "//app/admin/main/tag/model",
        "//app/admin/main/tag/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/permit",
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
