package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "http.go",
        "thumbup.go",
    ],
    importpath = "go-common/app/service/main/thumbup/server/http",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/thumbup/conf",
        "//app/service/main/thumbup/model",
        "//app/service/main/thumbup/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/middleware/rate",
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
