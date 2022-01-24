package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "up_test",
    srcs = [
        "api_test.go",
        "dao.cache_test.go",
        "dao_test.go",
        "mc.cache_test.go",
        "mc_all_identity_test.go",
        "mc_up_test.go",
        "sign_up_test.go",
        "up_info_test.go",
        "up_switch_test.go",
        "up_test.go",
    ],
    embed = [":up"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/up/conf",
        "//app/service/main/up/dao",
        "//app/service/main/up/dao/global",
        "//app/service/main/up/model",
        "//library/net/http/blademaster",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "up",
    srcs = [
        "api.go",
        "dao.cache.go",
        "dao.go",
        "mc.cache.go",
        "mc_all_identity.go",
        "mc_up.go",
        "sign_up.go",
        "up.go",
        "up_info.go",
        "up_switch.go",
    ],
    importpath = "go-common/app/service/main/up/dao/up",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/openplatform/article/model",
        "//app/interface/openplatform/article/rpc/client",
        "//app/service/main/up/api/v1:api",
        "//app/service/main/up/conf",
        "//app/service/main/up/dao/global",
        "//app/service/main/up/model",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
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
