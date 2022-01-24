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
    importpath = "go-common/app/interface/main/mcn/dao/global",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/mcn/conf",
        "//app/interface/main/tag/api",
        "//app/service/main/account/api",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/member/api",
        "//library/cache/memcache",
        "//library/log",
        "//library/net/http/blademaster",
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

go_test(
    name = "global_test",
    srcs = [
        "account_test.go",
        "resource_test.go",
    ],
    embed = [":global"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/mcn/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
