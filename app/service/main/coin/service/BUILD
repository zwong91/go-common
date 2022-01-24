package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "coin_test.go",
        "exp_test.go",
        "service_test.go",
        "usercoin_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/coin/api",
        "//app/service/main/coin/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "archive.go",
        "check.go",
        "coin.go",
        "exp.go",
        "service.go",
        "usercoin.go",
    ],
    importpath = "go-common/app/service/main/coin/service",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/coin/api",
        "//app/service/main/coin/conf",
        "//app/service/main/coin/dao",
        "//app/service/main/coin/model",
        "//app/service/main/member/api",
        "//app/service/main/member/model",
        "//app/service/main/member/model/block",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup.v2:errgroup_v2",
        "//library/sync/pipeline/fanout",
        "//library/time",
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
