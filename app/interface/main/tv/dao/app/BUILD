load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "app_test",
    srcs = [
        "dao_test.go",
        "homepage_test.go",
        "intervention_test.go",
        "module_test.go",
        "pgc_cards_test.go",
        "recommend_test.go",
        "redis_test.go",
        "region_test.go",
        "zone_test.go",
    ],
    embed = [":app"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/model",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "app",
    srcs = [
        "dao.go",
        "homepage.go",
        "intervention.go",
        "module.go",
        "pgc_cards.go",
        "recommend.go",
        "redis.go",
        "region.go",
        "zone.go",
    ],
    importpath = "go-common/app/interface/main/tv/dao/app",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/model",
        "//app/interface/main/tv/model/search",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
        "//library/xstr",
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
