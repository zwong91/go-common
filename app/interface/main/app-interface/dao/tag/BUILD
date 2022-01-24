package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "tag_test",
    srcs = ["dao_test.go"],
    embed = [":tag"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//app/interface/main/app-interface/model/tag",
        "//app/interface/main/tag/rpc/client",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "tag",
    srcs = ["dao.go"],
    importpath = "go-common/app/interface/main/app-interface/dao/tag",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//app/interface/main/app-interface/model/tag",
        "//app/interface/main/tag/model",
        "//app/interface/main/tag/rpc/client",
        "//library/ecode",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/xstr",
        "//vendor/github.com/pkg/errors",
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
