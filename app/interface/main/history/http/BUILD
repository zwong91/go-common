package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "history.go",
        "http.go",
        "param.go",
        "toview.go",
    ],
    importpath = "go-common/app/interface/main/history/http",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/history/conf",
        "//app/interface/main/history/model",
        "//app/interface/main/history/service",
        "//library/ecode",
        "//library/log",
        "//library/log/anticheat",
        "//library/log/infoc",
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
