package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "assist_test",
    srcs = ["service_test.go"],
    embed = [":assist"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/assist",
        "//app/interface/main/creative/service",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "assist",
    srcs = [
        "assist.go",
        "live.go",
        "revoc.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/creative/service/assist",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/account",
        "//app/interface/main/creative/dao/assist",
        "//app/interface/main/creative/dao/danmu",
        "//app/interface/main/creative/dao/reply",
        "//app/interface/main/creative/model/assist",
        "//app/interface/main/creative/service",
        "//app/service/main/account/model",
        "//library/ecode",
        "//library/log",
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
