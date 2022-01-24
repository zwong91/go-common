package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "appeal_test",
    srcs = ["service_test.go"],
    embed = [":appeal"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/appeal",
        "//app/interface/main/creative/service",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "appeal",
    srcs = [
        "appeal.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/creative/service/appeal",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/account",
        "//app/interface/main/creative/dao/appeal",
        "//app/interface/main/creative/dao/archive",
        "//app/interface/main/creative/dao/tag",
        "//app/interface/main/creative/model/appeal",
        "//app/interface/main/creative/model/archive",
        "//app/interface/main/creative/service",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//library/ecode",
        "//library/log",
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
