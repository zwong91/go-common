package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "block_test.go",
        "rule_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/spy/conf",
        "//app/service/main/spy/dao",
        "//app/service/main/spy/model",
        "//library/cache/redis",
        "//library/ecode",
        "//library/log",
        "//library/net/ip",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "account.go",
        "block.go",
        "rule.go",
        "service.go",
        "spy.go",
        "stat.go",
    ],
    importpath = "go-common/app/service/main/spy/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//app/service/main/spy/conf",
        "//app/service/main/spy/dao",
        "//app/service/main/spy/model",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/ip",
        "//library/net/metadata",
        "//library/stat/prom",
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
