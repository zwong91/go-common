package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "wechat_test",
    srcs = [
        "dao_test.go",
        "qrcode_test.go",
    ],
    embed = [":wechat"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/web-goblin/conf",
        "//app/interface/main/web-goblin/model/wechat",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "wechat",
    srcs = [
        "cache.go",
        "dao.cache.go",
        "dao.go",
        "qrcode.go",
    ],
    importpath = "go-common/app/interface/main/web-goblin/dao/wechat",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/web-goblin/conf",
        "//app/interface/main/web-goblin/model/wechat",
        "//library/cache",
        "//library/cache/redis",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/stat/prom",
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
