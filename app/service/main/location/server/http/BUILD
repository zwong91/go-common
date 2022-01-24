package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "http.go",
        "ip.go",
        "ip_old.go",
        "local.go",
        "manager.go",
        "zlimit.go",
    ],
    importpath = "go-common/app/service/main/location/server/http",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/location/conf",
        "//app/service/main/location/model",
        "//app/service/main/location/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/verify",
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
