package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "account_test",
    srcs = [
        "account_test.go",
        "bfs_test.go",
        "commercial_test.go",
        "official_test.go",
        "settings_test.go",
        "sms_test.go",
        "vip_test.go",
    ],
    embed = [":account"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/account/conf",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "account",
    srcs = [
        "account.go",
        "bfs.go",
        "commercial.go",
        "official.go",
        "settings.go",
        "sms.go",
        "vip.go",
    ],
    importpath = "go-common/app/interface/main/account/dao/account",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/account/conf",
        "//app/interface/main/account/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/xstr",
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
