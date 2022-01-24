load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "app_test",
    srcs = [
        "homepage_test.go",
        "module_test.go",
        "pgc_cards_test.go",
        "recommend_test.go",
        "region_test.go",
        "service_test.go",
    ],
    embed = [":app"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "app",
    srcs = [
        "auth.go",
        "follow.go",
        "homePage.go",
        "intervention.go",
        "module.go",
        "pages.go",
        "pgc_cards.go",
        "recommend.go",
        "region.go",
        "service.go",
        "tool.go",
        "ugc_types.go",
        "zone.go",
        "zone_index.go",
    ],
    importpath = "go-common/app/interface/main/tv/service/app",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/dao/app",
        "//app/interface/main/tv/dao/archive",
        "//app/interface/main/tv/dao/audit",
        "//app/interface/main/tv/dao/cms",
        "//app/interface/main/tv/dao/search",
        "//app/interface/main/tv/model",
        "//app/interface/main/tv/model/search",
        "//app/service/main/archive/api",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup",
        "//library/xstr",
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
