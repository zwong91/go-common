package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "feed",
    srcs = [
        "black.go",
        "cache.go",
        "index.go",
        "infoc.go",
        "operate.go",
        "player.go",
        "rank.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/app-intl/service/feed",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-card/model",
        "//app/interface/main/app-card/model/card",
        "//app/interface/main/app-card/model/card/ai",
        "//app/interface/main/app-card/model/card/banner",
        "//app/interface/main/app-card/model/card/cm",
        "//app/interface/main/app-card/model/card/operate",
        "//app/interface/main/app-card/model/card/rank",
        "//app/interface/main/app-intl/conf",
        "//app/interface/main/app-intl/dao/account",
        "//app/interface/main/app-intl/dao/archive",
        "//app/interface/main/app-intl/dao/black",
        "//app/interface/main/app-intl/dao/card",
        "//app/interface/main/app-intl/dao/location",
        "//app/interface/main/app-intl/dao/rank",
        "//app/interface/main/app-intl/dao/recommend",
        "//app/interface/main/app-intl/dao/relation",
        "//app/interface/main/app-intl/dao/tag",
        "//app/interface/main/app-intl/model",
        "//app/interface/main/app-intl/model/feed",
        "//app/interface/main/tag/model",
        "//app/service/main/account/model",
        "//app/service/main/archive/model/archive",
        "//app/service/main/location/model",
        "//app/service/main/relation/model",
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
