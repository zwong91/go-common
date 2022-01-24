package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.cache.go",
        "dao.go",
        "mc.go",
        "passport.go",
        "raw.go",
        "realname.go",
        "search.go",
        "vip.go",
    ],
    importpath = "go-common/app/service/main/account/dao",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/account/api",
        "//app/service/main/account/conf",
        "//app/service/main/account/model",
        "//app/service/main/member/api/gorpc",
        "//app/service/main/member/model",
        "//app/service/main/member/model/block",
        "//app/service/main/usersuit/model",
        "//app/service/main/usersuit/rpc/client",
        "//library/cache/memcache",
        "//library/database/elastic",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/stat/prom",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
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

go_test(
    name = "dao_test",
    srcs = [
        "dao.cache_test.go",
        "dao_test.go",
        "mc_test.go",
        "passport_test.go",
        "raw_test.go",
        "search_test.go",
        "vip_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/account/conf",
        "//app/service/main/account/model",
        "//app/service/main/member/model",
        "//app/service/main/member/model/block",
        "//library/cache/memcache",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
