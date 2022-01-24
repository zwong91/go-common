package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "account_test",
    srcs = ["dao_test.go"],
    embed = [":account"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/up/conf",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "account",
    srcs = [
        "account_cache.go",
        "dao.go",
        "rpc.go",
    ],
    importpath = "go-common/app/job/main/up/dao/account",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/up/conf",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//library/ecode",
        "//library/log",
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
