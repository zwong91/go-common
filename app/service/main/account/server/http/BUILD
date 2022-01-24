package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "acc.go",
        "http.go",
        "v1.go",
        "v2.go",
    ],
    importpath = "go-common/app/service/main/account/server/http",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/account/api",
        "//app/service/main/account/conf",
        "//app/service/main/account/model",
        "//app/service/main/account/service",
        "//app/service/main/member/model",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
