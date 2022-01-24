package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "account_test",
    srcs = ["service_test.go"],
    embed = [":account"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/account",
        "//app/interface/main/creative/service",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "account",
    srcs = [
        "account.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/creative/service/account",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/account",
        "//app/interface/main/creative/dao/archive",
        "//app/interface/main/creative/dao/article",
        "//app/interface/main/creative/dao/up",
        "//app/interface/main/creative/model/account",
        "//app/interface/main/creative/service",
        "//app/service/main/account/model",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/stat/prom",
        "//library/sync/errgroup",
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
