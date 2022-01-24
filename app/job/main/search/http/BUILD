package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "es.go",
        "http.go",
    ],
    importpath = "go-common/app/job/main/search/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/search/conf",
        "//app/job/main/search/model",
        "//app/job/main/search/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
