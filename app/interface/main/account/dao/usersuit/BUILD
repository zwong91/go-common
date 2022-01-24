package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "usersuit",
    srcs = [
        "api.go",
        "dao.go",
    ],
    importpath = "go-common/app/interface/main/account/dao/usersuit",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/account/conf",
        "//app/service/main/usersuit/model",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
    name = "usersuit_test",
    srcs = [
        "api_test.go",
        "dao_test.go",
    ],
    embed = [":usersuit"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/account/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
