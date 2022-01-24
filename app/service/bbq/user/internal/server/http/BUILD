package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "http.go",
        "user.go",
    ],
    importpath = "go-common/app/service/bbq/user/internal/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/user/api",
        "//app/service/bbq/user/internal/conf",
        "//app/service/bbq/user/internal/service",
        "//library/log",
        "//library/net/http/blademaster",
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
