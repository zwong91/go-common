package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "reply_test",
    srcs = ["service_test.go"],
    embed = [":reply"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/search",
        "//app/interface/main/creative/service",
        "//vendor/github.com/davecgh/go-spew/spew",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "reply",
    srcs = [
        "reply.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/creative/service/reply",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/account",
        "//app/interface/main/creative/dao/archive",
        "//app/interface/main/creative/dao/article",
        "//app/interface/main/creative/dao/elec",
        "//app/interface/main/creative/dao/music",
        "//app/interface/main/creative/dao/reply",
        "//app/interface/main/creative/dao/search",
        "//app/interface/main/creative/model/music",
        "//app/interface/main/creative/model/reply",
        "//app/interface/main/creative/model/search",
        "//app/interface/main/creative/service",
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup",
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
