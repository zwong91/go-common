package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "tag_test",
    srcs = ["tag_test.go"],
    embed = [":tag"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-tag/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "tag",
    srcs = [
        "api.go",
        "tag.go",
    ],
    importpath = "go-common/app/interface/main/app-tag/dao/tag",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-tag/conf",
        "//app/interface/main/app-tag/model/tag",
        "//app/interface/main/tag/model",
        "//app/interface/main/tag/rpc/client",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
