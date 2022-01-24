load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

package(default_visibility = ["//visibility:public"])

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

go_library(
    name = "http",
    srcs = [
        "app_auth.go",
        "config.go",
        "credential.go",
        "ecode.go",
        "http.go",
    ],
    importpath = "go-common/app/service/main/msm/http",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/msm/conf",
        "//app/service/main/msm/model",
        "//app/service/main/msm/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/verify",
    ],
)
