package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "weeklyhonor_test",
    srcs = ["service_test.go"],
    embed = [":weeklyhonor"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/up",
        "//app/interface/main/creative/dao/weeklyhonor",
        "//app/interface/main/creative/model/weeklyhonor",
        "//app/interface/main/creative/service",
        "//app/service/main/up/model",
        "//library/time",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "weeklyhonor",
    srcs = ["service.go"],
    importpath = "go-common/app/interface/main/creative/service/weeklyhonor",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/account",
        "//app/interface/main/creative/dao/archive",
        "//app/interface/main/creative/dao/up",
        "//app/interface/main/creative/dao/weeklyhonor",
        "//app/interface/main/creative/model/weeklyhonor",
        "//app/interface/main/creative/service",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/up/model",
        "//library/cache/memcache",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
        "//library/time",
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
