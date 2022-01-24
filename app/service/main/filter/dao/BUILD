package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "ai_test.go",
        "area_test.go",
        "dao_test.go",
        "filter_test.go",
        "hbase_test.go",
        "http_test.go",
        "key_test.go",
        "memcache_test.go",
        "rpc_test.go",
        "white_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/filter/conf",
        "//app/service/main/filter/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "ai.go",
        "area.go",
        "dao.go",
        "filter.go",
        "hbase.go",
        "http.go",
        "key.go",
        "memcache.go",
        "rpc.go",
        "white.go",
    ],
    importpath = "go-common/app/service/main/filter/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/antispam/model",
        "//app/service/main/antispam/rpc/client",
        "//app/service/main/filter/conf",
        "//app/service/main/filter/model",
        "//library/cache/memcache",
        "//library/database/hbase.v2:hbase_v2",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/time",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/tsuna/gohbase/hrpc",
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
