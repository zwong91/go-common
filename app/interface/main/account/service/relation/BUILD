package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "relation_test",
    srcs = ["relation_test.go"],
    embed = [":relation"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/account/conf",
        "//app/service/main/relation/model",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "relation",
    srcs = ["relation.go"],
    importpath = "go-common/app/interface/main/account/service/relation",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/account/conf",
        "//app/interface/main/account/dao/relation",
        "//app/interface/main/account/model",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/relation/model",
        "//app/service/main/relation/rpc/client",
        "//library/ecode",
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
