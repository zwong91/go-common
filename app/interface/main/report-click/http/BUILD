package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "click.go",
        "err_report.go",
        "http.go",
        "player.go",
    ],
    importpath = "go-common/app/interface/main/report-click/http",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/report-click/conf",
        "//app/interface/main/report-click/model",
        "//app/interface/main/report-click/service",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/verify",
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
