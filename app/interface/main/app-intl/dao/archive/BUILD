package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "archive_test",
    srcs = [
        "api_test.go",
        "dao_test.go",
        "memcache_test.go",
        "stat_test.go",
        "view_test.go",
    ],
    embed = [":archive"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-intl/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "archive",
    srcs = [
        "api.go",
        "dao.go",
        "memcache.go",
        "stat.go",
        "view.go",
    ],
    importpath = "go-common/app/interface/main/app-intl/dao/archive",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-intl/conf",
        "//app/interface/main/app-intl/model/player/archive",
        "//app/interface/main/app-intl/model/view",
        "//app/interface/main/history/model",
        "//app/interface/main/history/rpc/client",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//library/cache/memcache",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/sync/errgroup",
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
