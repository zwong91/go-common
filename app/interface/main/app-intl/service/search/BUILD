package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "search_test",
    srcs = ["search_test.go"],
    embed = [":search"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-intl/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "search",
    srcs = ["search.go"],
    importpath = "go-common/app/interface/main/app-intl/service/search",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-interface/model/tag",
        "//app/interface/main/app-intl/conf",
        "//app/interface/main/app-intl/dao/account",
        "//app/interface/main/app-intl/dao/archive",
        "//app/interface/main/app-intl/dao/article",
        "//app/interface/main/app-intl/dao/bangumi",
        "//app/interface/main/app-intl/dao/resource",
        "//app/interface/main/app-intl/dao/search",
        "//app/interface/main/app-intl/dao/tag",
        "//app/interface/main/app-intl/model",
        "//app/interface/main/app-intl/model/bangumi",
        "//app/interface/main/app-intl/model/search",
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//library/log",
        "//library/sync/errgroup",
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
