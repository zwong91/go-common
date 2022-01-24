package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "http_test",
    srcs = ["invite_test.go"],
    embed = [":http"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/smartystreets/goconvey/convey"],
)

go_library(
    name = "http",
    srcs = [
        "common.go",
        "http.go",
        "invite.go",
        "medal.go",
        "pendant.go",
        "upload.go",
    ],
    importpath = "go-common/app/admin/main/usersuit/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/usersuit/conf",
        "//app/admin/main/usersuit/model",
        "//app/admin/main/usersuit/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/permit",
        "//library/net/metadata",
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
