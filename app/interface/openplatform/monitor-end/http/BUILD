package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "alert.go",
        "http.go",
        "monitor.go",
    ],
    importpath = "go-common/app/interface/openplatform/monitor-end/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/openplatform/monitor-end/conf",
        "//app/interface/openplatform/monitor-end/model",
        "//app/interface/openplatform/monitor-end/model/monitor",
        "//app/interface/openplatform/monitor-end/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
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
