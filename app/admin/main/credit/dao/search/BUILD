package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "search_test",
    srcs = ["dao_test.go"],
    embed = [":search"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/credit/conf",
        "//app/admin/main/credit/model/blocked",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "search",
    srcs = ["dao.go"],
    importpath = "go-common/app/admin/main/credit/dao/search",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/credit/conf",
        "//app/admin/main/credit/model",
        "//app/admin/main/credit/model/blocked",
        "//app/admin/main/credit/model/search",
        "//library/database/elastic",
        "//library/log",
        "//library/net/metadata",
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
