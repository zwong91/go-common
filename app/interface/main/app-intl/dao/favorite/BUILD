package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "favorite_test",
    srcs = [
        "api_test.go",
        "dao_test.go",
    ],
    embed = [":favorite"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-intl/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "favorite",
    srcs = [
        "api.go",
        "dao.go",
    ],
    importpath = "go-common/app/interface/main/app-intl/dao/favorite",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-intl/conf",
        "//app/service/main/favorite/api/gorpc",
        "//app/service/main/favorite/model",
        "//library/ecode",
        "//library/net/http/blademaster",
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
