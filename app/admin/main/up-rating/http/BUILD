package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "authority.go",
        "csv.go",
        "http.go",
        "parameter.go",
        "score.go",
        "statistics.go",
    ],
    importpath = "go-common/app/admin/main/up-rating/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up-rating/conf",
        "//app/admin/main/up-rating/model",
        "//app/admin/main/up-rating/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/render",
        "//vendor/github.com/pkg/errors",
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
