package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "business.go",
        "challenge.go",
        "group.go",
        "http.go",
        "sign.go",
        "sobot.go",
        "tag.go",
    ],
    importpath = "go-common/app/service/main/workflow/http",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/workflow/conf",
        "//app/service/main/workflow/model",
        "//app/service/main/workflow/model/account",
        "//app/service/main/workflow/model/sobot",
        "//app/service/main/workflow/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
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
