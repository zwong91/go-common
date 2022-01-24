package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "account_test",
    srcs = ["mine_test.go"],
    embed = [":account"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "account",
    srcs = [
        "mine.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/app-interface/service/account",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//app/interface/main/app-interface/dao/account",
        "//app/interface/main/app-interface/dao/audio",
        "//app/interface/main/app-interface/dao/bplus",
        "//app/interface/main/app-interface/dao/member",
        "//app/interface/main/app-interface/dao/pay",
        "//app/interface/main/app-interface/dao/relation",
        "//app/interface/main/app-interface/dao/sidebar",
        "//app/interface/main/app-interface/model",
        "//app/interface/main/app-interface/model/space",
        "//app/service/main/account/model",
        "//app/service/main/relation/model",
        "//app/service/main/resource/model",
        "//library/log",
        "//library/sync/errgroup.v2:errgroup_v2",
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
