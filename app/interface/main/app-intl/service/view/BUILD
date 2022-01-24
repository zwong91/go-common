package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "view",
    srcs = [
        "access.go",
        "contain.go",
        "infoc.go",
        "service.go",
        "view.go",
    ],
    importpath = "go-common/app/interface/main/app-intl/service/view",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-intl/conf",
        "//app/interface/main/app-intl/dao/account",
        "//app/interface/main/app-intl/dao/archive",
        "//app/interface/main/app-intl/dao/assist",
        "//app/interface/main/app-intl/dao/audio",
        "//app/interface/main/app-intl/dao/bangumi",
        "//app/interface/main/app-intl/dao/coin",
        "//app/interface/main/app-intl/dao/dm",
        "//app/interface/main/app-intl/dao/favorite",
        "//app/interface/main/app-intl/dao/location",
        "//app/interface/main/app-intl/dao/manager",
        "//app/interface/main/app-intl/dao/region",
        "//app/interface/main/app-intl/dao/relation",
        "//app/interface/main/app-intl/dao/resource",
        "//app/interface/main/app-intl/dao/tag",
        "//app/interface/main/app-intl/dao/thumbup",
        "//app/interface/main/app-intl/dao/vip",
        "//app/interface/main/app-intl/model",
        "//app/interface/main/app-intl/model/bangumi",
        "//app/interface/main/app-intl/model/manager",
        "//app/interface/main/app-intl/model/region",
        "//app/interface/main/app-intl/model/view",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//app/service/main/location/model",
        "//app/service/main/thumbup/model",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//library/net/metadata",
        "//library/stat/prom",
        "//library/sync/errgroup",
        "//library/text/translate/chinese",
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
