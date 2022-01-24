package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "passport",
    srcs = ["passport.go"],
    importpath = "go-common/app/interface/main/account/service/passport",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/account/conf",
        "//app/interface/main/account/dao/passport",
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

go_test(
    name = "passport_test",
    srcs = ["passport_test.go"],
    embed = [":passport"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/account/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
