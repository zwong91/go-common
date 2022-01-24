package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = ["feed_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/web-feed/conf",
        "//app/interface/main/web-feed/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "feed.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/web-feed/service",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/web-feed/conf",
        "//app/interface/main/web-feed/dao",
        "//app/interface/main/web-feed/model",
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//app/service/main/feed/model",
        "//app/service/main/feed/rpc/client",
        "//library/cache",
        "//library/log",
        "//library/net/metadata",
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
