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
        "append_test.go",
        "dao_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/filter/conf",
        "//app/admin/main/filter/model",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "ai.go",
        "append.go",
        "area.go",
        "dao.go",
        "filter.go",
        "hbase.go",
        "key.go",
        "mc.go",
        "white.go",
    ],
    importpath = "go-common/app/admin/main/filter/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/filter/conf",
        "//app/admin/main/filter/model",
        "//library/cache/memcache",
        "//library/database/hbase.v2:hbase_v2",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/xstr",
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
    srcs = [
        ":package-srcs",
        "//app/admin/main/filter/dao/search:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
