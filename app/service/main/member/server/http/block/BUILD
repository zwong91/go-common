package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "block",
    srcs = [
        "block.go",
        "http.go",
    ],
    importpath = "go-common/app/service/main/member/server/http/block",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/member/model/block",
        "//app/service/main/member/service/block",
        "//library/ecode",
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
