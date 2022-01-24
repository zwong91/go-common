package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "apps.go",
        "auths.go",
        "business.go",
        "dataplatform.go",
        "http.go",
        "tasks.go",
        "upload.go",
    ],
    importpath = "go-common/app/admin/main/push/http",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/push/conf",
        "//app/admin/main/push/model",
        "//app/admin/main/push/service",
        "//app/service/main/push/model",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/permit",
        "//library/time",
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
