package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "captcha.go",
        "http.go",
        "render_image.go",
    ],
    importpath = "go-common/app/interface/main/captcha/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/captcha/conf",
        "//app/interface/main/captcha/service",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/rate",
        "//library/net/http/blademaster/middleware/verify",
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
