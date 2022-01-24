package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "black_test",
    srcs = ["dao_test.go"],
    embed = [":black"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-intl/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "black",
    srcs = [
        "api.go",
        "dao.go",
        "redis.go",
    ],
    importpath = "go-common/app/interface/main/app-intl/dao/black",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-intl/conf",
        "//library/cache/redis",
        "//library/log",
        "//library/net/http/blademaster",
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
