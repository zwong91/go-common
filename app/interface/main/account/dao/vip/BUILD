package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
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
    name = "vip_test",
    srcs = [
        "api_test.go",
        "dao_test.go",
        "oauth2_api_test.go",
        "oauth2_client_test.go",
        "vip_test.go",
    ],
    embed = [":vip"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/account/conf",
        "//app/interface/main/account/model",
        "//app/service/main/vip/model",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "vip",
    srcs = [
        "api.go",
        "dao.go",
        "oauth2_api.go",
        "oauth2_client.go",
    ],
    importpath = "go-common/app/interface/main/account/dao/vip",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/account/conf",
        "//app/interface/main/account/model",
        "//app/service/main/vip/model",
        "//library/conf/env",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//vendor/github.com/pkg/errors",
    ],
)
