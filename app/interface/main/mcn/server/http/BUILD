package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "http",
    srcs = [
        "api.go",
        "creative.go",
        "data.go",
        "http.go",
        "http_base.go",
        "mcn.go",
        "upload.go",
    ],
    importpath = "go-common/app/interface/main/mcn/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/mcn/conf",
        "//app/interface/main/mcn/model/mcnmodel",
        "//app/interface/main/mcn/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/antispam",
        "//library/net/http/blademaster/middleware/auth",
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

go_test(
    name = "http_test",
    srcs = ["http_test.go"],
    embed = [":http"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/mcn/model",
        "//library/net/http/blademaster/binding",
    ],
)
