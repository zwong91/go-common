package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "archive.go",
        "bws.go",
        "http.go",
        "kfc.go",
        "likes.go",
        "match.go",
        "subject.go",
    ],
    importpath = "go-common/app/admin/main/activity/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/activity/conf",
        "//app/admin/main/activity/model",
        "//app/admin/main/activity/model/kfc",
        "//app/admin/main/activity/service",
        "//app/admin/main/activity/service/kfc",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/permit",
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
