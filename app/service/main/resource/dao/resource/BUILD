package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "resource_test",
    srcs = [
        "banner_test.go",
        "dao_test.go",
        "resource_test.go",
    ],
    embed = [":resource"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/resource/conf",
        "//library/database/sql",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "resource",
    srcs = [
        "banner.go",
        "dao.go",
        "resource.go",
    ],
    importpath = "go-common/app/service/main/resource/dao/resource",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/resource/conf",
        "//app/service/main/resource/model",
        "//library/database/sql",
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
