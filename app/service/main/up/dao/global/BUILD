package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "global_test",
    srcs = ["resource_test.go"],
    embed = [":global"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/up/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "global",
    srcs = ["resource.go"],
    importpath = "go-common/app/service/main/up/dao/global",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/up/conf",
        "//library/database/sql",
        "//library/sync/pipeline/fanout",
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
