package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "base.go",
        "exp.go",
        "http.go",
        "moral.go",
        "realname.go",
    ],
    importpath = "go-common/app/service/main/member/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/member/api",
        "//app/service/main/member/conf",
        "//app/service/main/member/model",
        "//app/service/main/member/server/http/block",
        "//app/service/main/member/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/metadata",
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
    srcs = [
        ":package-srcs",
        "//app/service/main/member/server/http/block:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
