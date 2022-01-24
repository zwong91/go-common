package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "realname_test",
    srcs = ["realname_test.go"],
    embed = [":realname"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/account/conf",
        "//app/interface/main/account/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "realname",
    srcs = [
        "alipay.go",
        "antispam.go",
        "data.go",
        "geetest.go",
        "realname.go",
        "upload.go",
    ],
    importpath = "go-common/app/interface/main/account/service/realname",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/account/conf",
        "//app/interface/main/account/dao/realname",
        "//app/interface/main/account/model",
        "//app/interface/main/account/service/realname/crypto",
        "//app/service/main/member/api/gorpc",
        "//app/service/main/member/model",
        "//library/ecode",
        "//library/log",
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
    srcs = [
        ":package-srcs",
        "//app/interface/main/account/service/realname/crypto:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
