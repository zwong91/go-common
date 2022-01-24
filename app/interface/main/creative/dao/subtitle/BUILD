package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "subtitle",
    srcs = [
        "dao.go",
        "subtitle.go",
    ],
    importpath = "go-common/app/interface/main/creative/dao/subtitle",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/dm2/model",
        "//app/interface/main/dm2/rpc/client",
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

go_test(
    name = "subtitle_test",
    srcs = [
        "dao_test.go",
        "subtitle_test.go",
    ],
    embed = [":subtitle"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/dm2/model",
        "//app/interface/main/dm2/rpc/client",
        "//library/ecode",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
