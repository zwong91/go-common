package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "bnj",
    srcs = [
        "cache.go",
        "dao.go",
        "mc.cache.go",
        "push.go",
        "wechat.go",
    ],
    importpath = "go-common/app/job/main/activity/dao/bnj",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/activity/conf",
        "//library/cache/memcache",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
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

go_test(
    name = "bnj_test",
    srcs = [
        "cache_test.go",
        "dao_test.go",
        "mc.cache_test.go",
        "push_test.go",
        "wechat_test.go",
    ],
    embed = [":bnj"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/activity/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)
