package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "global",
    srcs = [
        "account.go",
        "resource.go",
    ],
    importpath = "go-common/app/admin/main/up-rating/dao/global",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up-rating/conf",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
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
    name = "global_test",
    srcs = [
        "account_test.go",
        "resource_test.go",
    ],
    embed = [":global"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/up-rating/conf",
        "//library/net/rpc",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
