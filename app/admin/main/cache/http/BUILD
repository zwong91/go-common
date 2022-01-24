package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "cache.go",
        "http.go",
        "overlord.go",
    ],
    importpath = "go-common/app/admin/main/cache/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/cache/conf",
        "//app/admin/main/cache/model",
        "//app/admin/main/cache/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/gopkg.in/yaml.v2:yaml_v2",
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
